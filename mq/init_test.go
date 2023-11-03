package mq_test

import (
	"github.com/ice-cream-heaven/log"
	"github.com/ice-cream-heaven/utils/common"
	"github.com/ice-cream-heaven/utils/mq"
	"sync"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	type TempData struct {
		Temp int
	}

	topic := mq.NewTopic[*TempData]("temp")

	var w sync.WaitGroup
	topic.AddChannel(&mq.Channel{}, func(msg *mq.Message[*TempData]) error {
		defer w.Done()
		log.Info(msg.Value.Temp)
		return nil
	})

	for i := 0; i < 1000; i++ {
		w.Add(1)
		err := topic.Push(&TempData{
			Temp: i,
		})
		if err != nil {
			log.Errorf("err:%v", err)
		}
	}

	w.Wait()
}

func TestRetry(t *testing.T) {
	type TempData struct {
		Temp int
	}

	topic := mq.NewTopic[*TempData]("temp")

	var w sync.WaitGroup
	w.Add(10)

	topic.AddChannel(&mq.Channel{}, func(msg *mq.Message[*TempData]) error {
		defer w.Done()

		return common.NewError(1).SetRetryDelay(time.Second).SetSkipRetryCount(true)
	})

	w.Add(10)
	err := topic.Push(&TempData{
		Temp: 1,
	})
	if err != nil {
		log.Errorf("err:%v", err)
	}

	w.Wait()
}
