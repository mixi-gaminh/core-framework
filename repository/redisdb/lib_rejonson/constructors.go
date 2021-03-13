package rejonson

import (
	redis "github.com/go-redis/redis/v8"
)

// ExtendClient - ExtendClient
func ExtendClient(client *redis.Client) *Client {
	return &Client{
		client,
		&redisProcessor{
			Process: client.Process,
		},
	}
}

// ExtendPipeline - ExtendPipeline
func ExtendPipeline(pipeline redis.Pipeliner) *Pipeline {
	return &Pipeline{
		pipeline,
		&redisProcessor{
			Process: pipeline.Process,
		},
	}
}
