package kafka

import (
	logger "github.com/mixi-gaminh/core-framework/logs"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// InitProducer - InitProducer
func (k *Kafka) InitProducer(kafkaURL string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaURL})
	if err != nil {
		return nil, err
	}
	return p, err
}

// CloseProducer - CloseProducer
func (k *Kafka) CloseProducer(producer *kafka.Producer) {
	logger.INFO("Closed Producer")
	defer producer.Close()
}

// PublishMessage - PublishMessage
func (k *Kafka) PublishMessage(p *kafka.Producer, topic string, partition int32, key, value string) error {
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
		Key:            []byte(key),
		Value:          []byte(value),
	}, nil)
	if err != nil {
		logger.ERROR(err)
		return err
	}
	p.Flush(10000)
	return nil
}
