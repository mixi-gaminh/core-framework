package kafka

import (
	logger "github.com/mixi-gaminh/core-framework/logs"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// KafkaPublishSave - KafkaPublishSave
func (k *Kafka) KafkaPublishSave(producer *kafka.Producer, collection, record, data string) error {
	keyMsgKafka := collection + "," + record
	valueMsgKafka := data
	if err := k.PublishMessage(producer, producerSaveCmdTopic, 0, keyMsgKafka, valueMsgKafka); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// KafkaPublishUpdate - KafkaPublishUpdate
func (k *Kafka) KafkaPublishUpdate(producer *kafka.Producer, collection, record, data string) error {
	keyMsgKafka := collection + "," + record
	valueMsgKafka := data
	if err := k.PublishMessage(producer, producerUpdateCmdTopic, 0, keyMsgKafka, valueMsgKafka); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// KafkaPublishDelete - KafkaPublishDelete
func (k *Kafka) KafkaPublishDelete(producer *kafka.Producer, collection, record string) error {
	keyMsgKafka := collection
	valueMsgKafka := record
	if err := k.PublishMessage(producer, producerDeleteCmdTopic, 0, keyMsgKafka, valueMsgKafka); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// KafkaPublishDrop - KafkaPublishDrop
func (k *Kafka) KafkaPublishDrop(producer *kafka.Producer, key, value string) error {
	keyMsgKafka := key
	valueMsgKafka := value
	if err := k.PublishMessage(producer, producerDropCmdTopic, 0, keyMsgKafka, valueMsgKafka); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// KafkaPublishES - KafkaPublishES
func (k *Kafka) KafkaPublishES(producer *kafka.Producer, topic, header, data string) error {
	consumerTopic := topic
	keyMsgKafka := header
	valueMsgKafka := data
	if err := k.PublishMessage(producer, consumerTopic, 0, keyMsgKafka, valueMsgKafka); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}
