package kafka

import (
	"context"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// IKafka - IKafka
type IKafka interface {
	KafkaConstructor()
	InitConsumer() (*kafka.Consumer, error)
	CloseConsumer(c *kafka.Consumer)

	// SubscribeTopic(c *kafka.Consumer, topic []string)
	SubscribeTopic(*kafka.Consumer, []string)

	// ReceiveMessage(c *kafka.Consumer, callbackFunction callbackFn)
	ReceiveMessage(*kafka.Consumer, callbackFn)

	// ReceiveMessageQueue(ctx context.Context, c *kafka.Consumer, callbackFunction callbackFnMQ)
	ReceiveMessageQueue(context.Context, *kafka.Consumer, callbackFnMQ)

	// ReceiveMessageQueueSyncDB(ctx context.Context, c *kafka.Consumer, callbackFunction callbackFnMQ)
	ReceiveMessageQueueSyncDB(context.Context, *kafka.Consumer, callbackFnMQ)

	// KafkaPublishSave(producer *kafka.Producer, collection, record, data string) error
	KafkaPublishSave(*kafka.Producer, string, string, string) error

	// KafkaPublishUpdate(producer *kafka.Producer, collection, record, data string) error
	KafkaPublishUpdate(*kafka.Producer, string, string, string) error

	// KafkaPublishDelete(producer *kafka.Producer, collection, record string) error
	KafkaPublishDelete(*kafka.Producer, string, string) error
}
