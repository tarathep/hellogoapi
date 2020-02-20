package event

import (
	"github.com/Shopify/sarama"
)

type SyncProducer struct {
	sarama.SyncProducer
}

func Init(brokerAddress string) *SyncProducer {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	//config.Producer.Retry.Max = 5
	conf.Producer.Return.Successes = true

	brokers := []string{brokerAddress}
	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		panic(err)
	}
	return &SyncProducer{producer}
}
