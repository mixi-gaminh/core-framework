package logs

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nxadm/tail"
	"github.com/spf13/viper"
)

// StreamToNats - StreamToNats
func StreamToNats() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if viper.GetString(`nats.active`) != "true" {
		return
	}
	logNameSubject := viper.GetString(`context.name`) + "__logging"
	natsURL := viper.GetString(`nats.url`)
	go func(natsURL, logNameSubject string) {
		opts := []nats.Option{nats.Name("LogStreaming")}
		nc, err := nats.Connect(natsURL, opts...)
		if err != nil {
			log.Fatal(err)
		}

		// STREAM LOG INFO
		infoLog, err := tail.TailFile(INFO_LOG_FILE, tail.Config{Follow: true})
		if err != nil {
			panic(err)
		}
		for line := range infoLog.Lines {
			if line.Text != "" {
				nc.Publish(logNameSubject, []byte(line.Text))
			}
		}

		// STREAM LOG ERROR
		errorLog, err := tail.TailFile(ERROR_LOG_FILE, tail.Config{Follow: true})
		if err != nil {
			panic(err)
		}
		for line := range errorLog.Lines {
			if line.Text != "" {
				nc.Publish(logNameSubject, []byte(line.Text))
			}
		}

		// STREAM LOG WARNING
		warningLog, err := tail.TailFile(WARNING_LOG_FILE, tail.Config{Follow: true})
		if err != nil {
			panic(err)
		}
		for line := range warningLog.Lines {
			if line.Text != "" {
				nc.Publish(logNameSubject, []byte(line.Text))
			}
		}
	}(natsURL, logNameSubject)
}

// ListenLogNats - ListenLogNats
func ListenLogNats() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if viper.GetString(`nats.active`) != "true" {
		return
	}
	logNameSubject := viper.GetString(`context.name`) + "__logging"
	natsURL := viper.GetString(`nats.url`)
	queueName := viper.GetString(`nats.queue_name`)
	go func(natsURL, queueName, logNameSubject string) {
		opts := []nats.Option{nats.Name("LogStreaming")}
		nc, err := nats.Connect(natsURL, opts...)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := nc.QueueSubscribe(logNameSubject, queueName, logStreamHandler); err != nil {
			log.Fatal(err)
		}
	}(natsURL, queueName, logNameSubject)
}

func logStreamHandler(msg *nats.Msg) {
	fmt.Println(string(msg.Data))
}
