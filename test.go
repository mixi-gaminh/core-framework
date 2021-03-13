package main

import (
	logger "github.com/mixi-gaminh/core-framework/logs"
	kafkaMQ "github.com/mixi-gaminh/core-framework/messages/kafka"
	natsMQ "github.com/mixi-gaminh/core-framework/messages/nats"
	"github.com/mixi-gaminh/core-framework/queue"
	miniodb "github.com/mixi-gaminh/core-framework/repository/miniodb"
	mongodb "github.com/mixi-gaminh/core-framework/repository/mongodb"
	redisdb "github.com/mixi-gaminh/core-framework/repository/redisdb"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

var redis redisdb.Cache
var mgodb mongodb.Mgo
var minio miniodb.FileStorage
var k kafkaMQ.Kafka
var natsSingle, natsStream natsMQ.NATS
var q queue.Queue

func main() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	VDDDataBaseName := viper.GetString(`mongoselectSession().DBname`)

	// initialize kafka constructor
	k.KafkaConstructor(
		viper.GetString(`kafka.active`),
		viper.GetString(`kafka.addr`),
		viper.GetString(`kafka.groupID`),
		viper.GetString(`kafka.autoOffsetReset`),
		viper.GetString(`kafka.maxPollIntervalms`),
		viper.GetString(`kafka.sessionTimeoutms`),
		viper.GetString(`kafka.messageMaxBytes`))
	// initialize redis constructor
	redis.RedisConstructor(
		viper.GetString(`redis.url`),
		viper.GetInt(`redis.max_clients`),
		viper.GetInt(`redis.min_idle`))
	// initialize mongo constructor
	mgodb.MongoDBConstructor(viper.GetStringSlice(`mongodb.url`),
		viper.GetString(`mongodb.username`),
		viper.GetString(`mongodb.password`))
	// initialize queueconstructor constructor
	q.QueueConstructor(viper.GetString(`centrifugoWS.url`), VDDDataBaseName)
	// initialize minio constructor
	minio.FileStorageConstructor(
		viper.GetString(`minio.host`),
		viper.GetString(`minio.minioEndpoint`),
		viper.GetString(`minio.minioAccessKeyID`),
		viper.GetString(`minio.minioSecretAccessKey`),
		viper.GetString(`minio.minioLocation`),
		viper.GetString(`domain.url`)+"/"+viper.GetString(`domain.storage_prefix`),
		viper.GetBool(`minio.minioUseSSL`))
	// initialize nats single constructor
	natsSingle.NATSConstructor(
		"test",
		"SINGLE",
		viper.GetString(`nats.active`),
		viper.GetString(`nats.url`),
		viper.GetString(`nats.queue_name`),
		viper.GetString(`nats.request_subject`),
		viper.GetString(`nats.response_subject`),
		viper.GetString(`nats.queue_name_stream`),
		viper.GetString(`nats.request_stream_subject`),
		viper.GetString(`nats.response_stream_subject`), test)
	// initialize nats stream constructor
	natsStream.NATSConstructor(
		"test",
		"STREAM",
		viper.GetString(`nats.active`),
		viper.GetString(`nats.url`),
		viper.GetString(`nats.queue_name`),
		viper.GetString(`nats.request_subject`),
		viper.GetString(`nats.response_subject`),
		viper.GetString(`nats.queue_name_stream`),
		viper.GetString(`nats.request_stream_subject`),
		viper.GetString(`nats.response_stream_subject`), test)

	logger.Constructor()
	logger.INFO("Test Lib OK")
	defer logger.Close()
}

func test(msg *nats.Msg) {}
