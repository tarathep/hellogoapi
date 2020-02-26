package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type SyncProducer struct {
	sarama.SyncProducer
}

func InitProducer(brokerAddress string) *SyncProducer {
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

type Producer interface {
	Publish(topic string, message string) error
}

func (produce SyncProducer) Publish(topic string, message string) error {

	partition, offset, err := produce.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
