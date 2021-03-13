package kafka

import (
	"context"
	"log"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type callbackFn func([]string)
type callbackFnMQ func(context.Context, []string)

func callbackFnHandler(f callbackFn, val []string) {
	f(val)
}

func callbackFnMQHandler(ctx context.Context, f callbackFnMQ, val []string) {
	f(ctx, val)
}

// InitConsumer - InitConsumer
func (k *Kafka) InitConsumer() (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    k.KafkaURL,
		"group.id":             k.KafkaGroupID,
		"auto.offset.reset":    k.KafkaAutoOffsetReset,
		"max.poll.interval.ms": k.KafkaMaxPollIntervalms,
		"session.timeout.ms":   k.KafkaSessionTimeoutms,
		"message.max.bytes":    k.KafkaMessageMaxBytes,
	})

	if err != nil {
		return nil, err
	}
	return c, err
}

// CloseConsumer - CloseConsumer
func (k *Kafka) CloseConsumer(c *kafka.Consumer) {
	//logger.INFO("Closed Consumer")
	c.Close()
}

// SubscribeTopic - SubscribeTopic
func (k *Kafka) SubscribeTopic(c *kafka.Consumer, topic []string) {
	c.SubscribeTopics(topic, nil)
}

// ReceiveMessage - ReceiveMessage
func (k *Kafka) ReceiveMessage(c *kafka.Consumer, callbackFunction callbackFn) {
	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				//log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				keyMsg := string(msg.Key)
				valueMsg := string(msg.Value)

				paramCallbackFn := []string{keyMsg, valueMsg}
				callbackFnHandler(callbackFunction, paramCallbackFn)
			} else {
				// The client will automatically try to recover from all errors.
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
}

// ReceiveMessageQueue - ReceiveMessageQueue
func (k *Kafka) ReceiveMessageQueue(ctx context.Context, c *kafka.Consumer, callbackFunction callbackFnMQ) {
	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				//log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				keyMsg := string(msg.Key)
				valueMsg := string(msg.Value)

				paramCallbackFn := []string{keyMsg, valueMsg}
				callbackFnMQHandler(ctx, callbackFunction, paramCallbackFn)
			} else {
				// The client will automatically try to recover from all errors.
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
}

// ReceiveMessageQueueSyncDB - ReceiveMessageQueueSyncDB
func (k *Kafka) ReceiveMessageQueueSyncDB(ctx context.Context, c *kafka.Consumer, callbackFunction callbackFnMQ) {
	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				keyMsg := string(msg.Key)
				valueMsg := string(msg.Value)

				paramCallbackFn := []string{keyMsg, valueMsg}
				callbackFnMQHandler(ctx, callbackFunction, paramCallbackFn)
			} else {
				// The client will automatically try to recover from all errors.
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
}
