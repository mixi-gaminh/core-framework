package kafka

import logger "github.com/mixi-gaminh/core-framework/logs"

// KafkaMQ - KafkaMQ
var KafkaMQ Kafka

// Kafka - Kafka
type Kafka struct {
	KafkaURL               string
	KafkaGroupID           string
	KafkaAutoOffsetReset   string
	KafkaMaxPollIntervalms string
	KafkaSessionTimeoutms  string
	KafkaMessageMaxBytes   string
	ProducerSaveTopic      string
	ProducerUpdateTopic    string
	ProducerDeleteTopic    string
	ProducerDropTopic      string
}

// KafkaConstructor - KafkaConstructor
func (k *Kafka) KafkaConstructor(_active, _kafkaURL, _afkaGroupID, _kafkaAutoOffsetReset,
	_kafkaMaxPollIntervalms, _kafkaSessionTimeoutms, _kafkaMessageMaxBytes,
	_producerSaveTopic, _producerUpdateTopic, _producerDeleteTopic, _producerDropTopic string) {
	if _active == "true" {
		k.KafkaURL = _kafkaURL
		k.KafkaGroupID = _afkaGroupID
		k.KafkaAutoOffsetReset = _kafkaAutoOffsetReset
		k.KafkaMaxPollIntervalms = _kafkaMaxPollIntervalms
		k.KafkaSessionTimeoutms = _kafkaSessionTimeoutms
		k.KafkaMessageMaxBytes = _kafkaMessageMaxBytes
		k.ProducerSaveTopic = _producerSaveTopic
		k.ProducerUpdateTopic = _producerUpdateTopic
		k.ProducerDeleteTopic = _producerDeleteTopic
		k.ProducerDropTopic = _producerDropTopic
	}
	//logger.Constructor(logger.IsDevelopment)
	logger.NewLogger()
	logger.INFO("Kafka Constructor Successfull")
}
