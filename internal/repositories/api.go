package repositories

import (
	"context"

	"go.zenithar.org/pubsubhub/internal/models"
)

// EventStore is the event storage contract
type EventStore interface {
	Create(ctx context.Context, entity *models.Event) error
}
