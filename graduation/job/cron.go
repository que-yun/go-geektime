package job

import (
	"github.com/go-kratos/kratos/v2/config"
	"go-geektime/graduation/internal/conf"
	"go-geektime/graduation/pkg/mq"
)

// 定时任务

func NewJobService(c config.Config, conf *conf.Data) *mq.SaramaConsumerJobServer {
	consumerGroup, err := mq.NewKafkaConsumeGroup(conf)
	if err != nil {
		return nil
	}
	topic, err := c.Value("topic").String()
	if err != nil {
		return nil
	}
	return mq.NewSaramaConsumerJobServer([]string{topic}, nil, consumerGroup)
}
