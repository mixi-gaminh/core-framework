package nats

import (
	"log"
	"strings"
	"time"

	logger "github.com/mixi-gaminh/core-framework/logs"
	"github.com/nats-io/nats.go"
)

var natsPoolClient [100]*nats.Conn
var err error

// NATSConstructor - NATSConstructor
func (n *NATS) NATSConstructor(_memberID, _typeReqRepl, _active, _natsURL, _queueName, _requestSubject,
	_responseSubject, _queueNameStream, _requestStreamSubject, _responseStreamSubject string, cb nats.MsgHandler) {
	if _active == "true" {
		n.NATSURL = _natsURL

		n.QueueName = _queueName
		n.RequestSubject = _requestSubject
		n.ResponseSubject = _responseSubject

		n.QueueNameStream = _queueNameStream
		n.RequestStreamSubject = _requestStreamSubject
		n.ResponseStreamSubject = _responseStreamSubject

		opts := []nats.Option{nats.Name("VDD_Request_Reply_" + _memberID)}
		opts = n.SetupConnOptions(opts)
		for _, nc := range natsPoolClient {
			// Connect to NATS
			nc, err = nats.Connect(n.NATSURL, opts...)
			if err != nil {
				logger.FATAL(err)
			}
			if strings.ToUpper(_typeReqRepl) == "STREAM" {
				if _, err := nc.QueueSubscribe(n.RequestStreamSubject, n.QueueNameStream, cb); err != nil {
					logger.FATAL(err)
				}
			} else if strings.ToUpper(_typeReqRepl) == "SINGLE" {
				if _, err := nc.QueueSubscribe(n.RequestSubject, n.QueueName, cb); err != nil {
					logger.FATAL(err)
				}
			} else {
				logger.FATAL("Type Request - Reply is invalid")
			}
		}
	}
	//logger.Constructor(logger.IsDevelopment)
	logger.NewLogger()
	logger.INFO("NATS Constructor Successfull")
}

// SetupConnOptions - SetupConnOptions
func (n *NATS) SetupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
