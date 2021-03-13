package queue

import (
	"context"

	"github.com/centrifugal/centrifuge-go"
)

// IQueue - IQueue
type IQueue interface {
	// OnConnect(c *centrifuge.Client, e centrifuge.ConnectEvent)
	OnConnect(*centrifuge.Client, centrifuge.ConnectEvent)

	// OnDisconnect(c *centrifuge.Client, e centrifuge.DisconnectEvent)
	OnDisconnect(*centrifuge.Client, centrifuge.DisconnectEvent)

	// CreateConnectionToCentrifugo() (*centrifuge.Client, error)
	CreateConnectionToCentrifugo() (*centrifuge.Client, error)

	// Save(ctx context.Context, msg []string)
	Save(context.Context, []string)

	// Update(ctx context.Context, msg []string)
	Update(context.Context, []string)

	// Delete(ctx context.Context, msg []string)
	Delete(context.Context, []string)

	// Drop(ctx context.Context, msg []string)
	Drop(context.Context, []string)

	// SynchronizeEvent(ctx context.Context, memberID string)
	SynchronizeEvent(context.Context, string)
}
