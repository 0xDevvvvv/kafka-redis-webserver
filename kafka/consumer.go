package kafka

import (
	"context"
	"log"

	"github.com/0xDevvvvv/kafka-redis-webserver/redis"
	"github.com/IBM/sarama"
)

type Consumer struct{}

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session
func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		key := string(message.Key)
		value := string(message.Value)

		log.Printf("🔄 Consumed message: key=%s, value=%s", key, value)
		err := redis.Set(context.Background(), key, value)
		if err != nil {
			log.Printf("Failed to write to Redis: %v\n", err)
		}
		session.MarkMessage(message, "")
	}
	return nil
}

func StartConsumer(brokers []string, topic, groupID string) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("❌ Failed to create consumer group: %v", err)
	}

	go func() {
		for {
			if err := consumerGroup.Consume(context.Background(), []string{topic}, &Consumer{}); err != nil {
				log.Printf("⚠️ Kafka consume error: %v", err)
			}
		}
	}()

}
