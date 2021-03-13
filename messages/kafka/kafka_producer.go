package kafka

import (
	"log"

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
	log.Println("Closed Producer")
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
		log.Println(err)
		return err
	}
	p.Flush(10000)
	return nil
}
