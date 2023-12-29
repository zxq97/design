package kafka

import (
	"github.com/Shopify/sarama"
)

func NewSyncProducer(addr []string) (sarama.SyncProducer, error) {
	kfkConf := sarama.NewConfig()
	kfkConf.Producer.RequiredAcks = sarama.WaitForAll
	kfkConf.Producer.Retry.Max = 3
	kfkConf.Producer.Return.Successes = true
	return sarama.NewSyncProducer(addr, kfkConf)
}

func NewConsumerGroup(addr []string, group string) (sarama.ConsumerGroup, error) {
	kfkConf := sarama.NewConfig()
	kfkConf.Consumer.Return.Errors = true
	kfkConf.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	kfkConf.Consumer.Offsets.Initial = sarama.OffsetNewest
	return sarama.NewConsumerGroup(addr, group, kfkConf)
}
