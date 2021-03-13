package kafka

import logger "github.com/mixi-gaminh/core-framework/logs"

// KafkaMQ - KafkaMQ
var KafkaMQ Kafka

var producerSaveCmdTopic string = "SaveRecord"
var producerUpdateCmdTopic string = "UpdateRecord"
var producerDeleteCmdTopic string = "DeleteRecord"
var producerDropCmdTopic string = "Drop"

// Kafka - Kafka
type Kafka struct {
	KafkaURL               string
	KafkaGroupID           string
	KafkaAutoOffsetReset   string
	KafkaMaxPollIntervalms string
	KafkaSessionTimeoutms  string
	KafkaMessageMaxBytes   string
}

// KafkaConstructor - KafkaConstructor
func (k *Kafka) KafkaConstructor(_active, _kafkaURL, _afkaGroupID, _kafkaAutoOffsetReset, _kafkaMaxPollIntervalms, _kafkaSessionTimeoutms, _kafkaMessageMaxBytes string) {
	if _active == "true" {
		k.KafkaURL = _kafkaURL
		k.KafkaGroupID = _afkaGroupID
		k.KafkaAutoOffsetReset = _kafkaAutoOffsetReset
		k.KafkaMaxPollIntervalms = _kafkaMaxPollIntervalms
		k.KafkaSessionTimeoutms = _kafkaSessionTimeoutms
		k.KafkaMessageMaxBytes = _kafkaMessageMaxBytes
	}
	logger.Constructor()
	logger.NewLogger()
	logger.INFO("Kafka Constructor Successfull")
}
