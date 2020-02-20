package event

import (
	"fmt"

	"github.com/Shopify/sarama"
)

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
