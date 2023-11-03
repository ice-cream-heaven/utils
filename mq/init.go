package mq

import (
	"github.com/ice-cream-heaven/log"
	"github.com/ice-cream-heaven/utils/osx"
	"github.com/ice-cream-heaven/utils/runtime"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/nsqio/nsq/nsqd"
)

var (
	producer     *nsqd.NSQD
	producerOnce sync.Once
)

func producerLoad() {
	producerOnce.Do(func() {
		opt := nsqd.NewOptions()
		opt.HTTPAddress = ""
		opt.HTTPSAddress = ""
		opt.TCPAddress = ""
		opt.BroadcastAddress = ""
		opt.MaxMsgSize = 0
		opt.MemQueueSize = 0

		strconv.Atoi("")

		opt.DataPath = DataPath()

		log.Infof("mq data path:%s", opt.DataPath)

		if !osx.Exists(opt.DataPath) {
			err := os.MkdirAll(opt.DataPath, os.ModePerm)
			if err != nil {
				log.Panicf("err:%v", err)
				return
			}
		}

		opt.MaxMsgSize = 1024 * 1024 * 1024
		opt.Logger = NewLogger()

		var err error
		producer, err = nsqd.New(opt)
		if err != nil {
			log.Panicf("err:%v", err)
			return
		}

		go func() {
			err := producer.Main()
			if err != nil {
				log.Panicf("err:%v", err)
			}
		}()

		go func(c chan os.Signal) {
			for {
				select {
				case <-c:
					Close()

					return
				case <-time.After(time.Minute * 5):
					err := filepath.Walk(opt.DataPath, func(path string, info fs.FileInfo, err error) error {
						if info == nil {
							return nil
						}

						if info.IsDir() {
							return nil
						}

						switch filepath.Ext(path) {
						case ".bad":
							log.Warnf("remove bad file:%s", path)
							err := os.RemoveAll(path)
							if err != nil {
								log.Errorf("err:%v", err)
							}
						case ".tmp":
							if time.Now().AddDate(0, 0, -1).After(info.ModTime()) {
								log.Warnf("remove bad file:%s", path)
								err := os.RemoveAll(path)
								if err != nil {
									log.Errorf("err:%v", err)
								}
							}
						}
						return nil
					})
					if err != nil {
						log.Errorf("err:%v", err)
					}
				}
			}
		}(runtime.GetExitSign())
	})
}

func Close() {
	if producer == nil {
		return
	}

	for topicName, topic := range producer.CloneTopic() {
		for channelName, channel := range topic.CloneChannel() {
			err := channel.Close()
			if err != nil {
				log.Errorf("close channel %s-%s err:%v", topicName, channelName, err)
			}
		}

		err := topic.Close()
		if err != nil {
			log.Errorf("close topic %s err:%v", topicName, err)
		}
	}
}
