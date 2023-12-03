package mq

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/ice-cream-heaven/log"
	"github.com/ice-cream-heaven/utils/common"
	"github.com/ice-cream-heaven/utils/json"
	"github.com/nsqio/nsq/nsqd"
	"sync"
	"time"
)

type Topic[M any] struct {
	once  sync.Once
	name  string
	topic *nsqd.Topic
}

func NewTopic[M any](name string) *Topic[M] {
	p := &Topic[M]{
		name: name,
	}

	return p
}

type Channel struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	ConcurrentCount int `json:"concurrent_count,omitempty" yaml:"concurrent_count,omitempty" default:"10" help:"并发数"`

	MaxRetryCount uint16 `json:"max_retry_count,omitempty" yaml:"max_retry_count,omitempty" default:"0" help:"最大重试次数，如果小于等于0，就无限重试"`

	MaxExecTime time.Duration `json:"max_exec_time,omitempty" yaml:"max_exec_time,omitempty" default:"5s" help:"最大执行时间，执行超时时间"`

	MessageTimeout time.Duration `json:"message_timeout,omitempty" yaml:"message_timeout,omitempty" help:"消息超时时间，如果小于等于0标识没有超时时间"`
}

func (opt *Channel) apply() {
	if opt.Name == "" {
		opt.Name = "default"
	}

	if opt.ConcurrentCount <= 0 {
		opt.ConcurrentCount = 1
	}
}

func (p *Topic[M]) getTopic() *nsqd.Topic {
	p.once.Do(func() {
		producerLoad()

		p.topic = producer.GetTopic(p.name)
	})
	return p.topic
}

func (p *Topic[M]) AddChannel(opt *Channel, logic func(msg *Message[M]) error) {
	producerLoad()

	opt.apply()
	name := opt.Name

	if p.getTopic().ChannelExist(name) {
		log.Panicf("channel %s in %s already exist", name, p.name)
	}

	if opt.MaxExecTime <= 0 {
		opt.MaxExecTime = time.Minute * 10
	}

	channel := p.getTopic().GetChannel(name)

	handlerMsg := func(clientId int64, m *nsqd.Message) error {
		var msg Message[M]
		var err error

		log.SetTrace(fmt.Sprintf("%s.%s.%s", p.name, name, m.ID))
		defer log.DelTrace()

		log.Infof("got msg from %s.%s.%s", p.name, name, m.ID)

		m.Attempts++

		if opt.MaxRetryCount > 0 && m.Attempts > opt.MaxRetryCount {
			log.Warnf("msg %s.%s.%s attempts %d > %d", p.name, name, m.ID, m.Attempts, opt.MaxRetryCount)
			return channel.FinishMessage(clientId, m.ID)
		}

		err = channel.StartInFlightTimeout(m, clientId, opt.MaxExecTime)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil
		}

		err = json.Unmarshal(m.Body, &msg)
		if err != nil {
			log.Errorf("err:%v", err)
			return channel.FinishMessage(clientId, m.ID)
		}

		if opt.MessageTimeout > 0 && time.Now().Add(-opt.MessageTimeout).After(msg.CreatedAt) {
			log.Warnf("%s.%s.%s timeout", p.name, name, m.ID)
			return channel.FinishMessage(clientId, m.ID)
		}

		err = logic(&msg)
		if err != nil {
			log.Errorf("err:%v", err)

			if x, ok := err.(*common.Error); ok {
				if x.SkipRetryCount {
					m.Attempts--
				}

				if x.NeedRetry() || x.Code == common.StatusTooManyRequests {
					if x.RetryDelay > 0 {
						log.Warnf("%s.%s.%s retry delay:%v", p.name, name, m.ID, x.RetryDelay)

						return channel.RequeueMessage(clientId, m.ID, x.RetryDelay)
					} else {
						log.Warnf("%s.%s.%s retry delay:5s", p.name, name, m.ID)
						return channel.RequeueMessage(clientId, m.ID, time.Second*5)
					}
				}
			} else {
				// 系统错误全部都重试
				log.Warnf("%s.%s.%s retry delay:5s", p.name, name, m.ID)

				return channel.RequeueMessage(clientId, m.ID, time.Second*5)
			}

			return err
		}

		log.Infof("handler %s.%s.%s success", p.name, name, m.ID)

		return channel.FinishMessage(clientId, m.ID)
	}

	handler := func(clientId int64) {
		for {
			err := handlerMsg(clientId, channel.ReadMessage())
			if err != nil {
				log.Errorf("err:%v", err)
			}
		}
	}

	for i := 0; i < (opt.ConcurrentCount); i++ {
		go handler(int64(i))
	}
}

func (p *Topic[M]) PushMsg(msg *Message[M]) error {
	producerLoad()

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	log.Infof("push %s msg", p.name)

	return p.getTopic().PutMessage(nsqd.NewMessage(p.getTopic().GenerateID(), body))
}

func (p *Topic[M]) Push(m M) error {
	return p.PushMsg(&Message[M]{
		Value:     m,
		CreatedAt: time.Now(),
	})
}

func (p *Topic[M]) MultiPublishMsg(list []*Message[M]) error {
	if len(list) == 0 {
		return nil
	}

	producerLoad()

	log.Infof("push %s %d msg", p.name, len(list))

	return p.getTopic().PutMessages(pie.Map(list, func(m *Message[M]) *nsqd.Message {
		body, err := json.Marshal(m)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil
		}
		return nsqd.NewMessage(p.getTopic().GenerateID(), body)
	}))
}

func (p *Topic[M]) MultiPublish(list []M) error {
	if len(list) == 0 {
		return nil
	}

	return p.MultiPublishMsg(pie.Map(list, func(m M) *Message[M] {
		return &Message[M]{
			Value:     m,
			CreatedAt: time.Now(),
		}
	}))
}

func (p *Topic[M]) DeferredPublishMsg(delay time.Duration, msg *Message[M]) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	producerLoad()

	log.Infof("push %s msg", p.name)

	return p.getTopic().PutMessage(nsqd.NewMessage(p.getTopic().GenerateID(), body))
}

func (p *Topic[M]) DeferredPublish(delay time.Duration, m M) error {
	return p.DeferredPublishMsg(delay, &Message[M]{
		Value:     m,
		CreatedAt: time.Now(),
	})
}

func (p *Topic[M]) Depth(name string) int64 {
	producerLoad()

	var depth int64
	for _, channel := range p.getTopic().CloneChannel() {
		depth += channel.Depth() + channel.InFlightDepth() + channel.DeferredDepth()
	}

	return depth
}

func CloneTopic() map[string]*nsqd.Topic {
	producerLoad()

	return producer.CloneTopic()
}

func GetDepth() (depth int64) {
	producerLoad()

	for _, topic := range CloneTopic() {
		for _, channel := range topic.CloneChannel() {
			depth += channel.Depth() + channel.InFlightDepth() + channel.DeferredDepth()
		}
	}
	return
}
