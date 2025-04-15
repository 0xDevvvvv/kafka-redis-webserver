package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.AsyncProducer

func InitProducer(brokers []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3
	config.Producer.Partitioner = sarama.NewHashPartitioner

	var err error
	producer, err = sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error in initializing Kafka producer ", err)
	}

	go func() {
		for err := range producer.Errors() {
			log.Printf("⚠️ Kafka producer error: %v", err)
		}
	}()
}

func ProduceMessage(topic, key, value string) {
	go func() {
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(value),
		}
	}()
}
