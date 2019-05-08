package postgresql

import (
	"context"

	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/pubsubhub/internal/models"
	"go.zenithar.org/pubsubhub/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgEventStore struct {
	adapter *db.Default
}

// NewEventStore returns an PostgreSQL event store implementation
func NewEventStore(cfg *db.Configuration, session *sqlx.DB) repositories.EventStore {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "realm_id", "event_id", "type", "body", "inserted_at",
	}

	// Sortable columns
	sortableColumns := []string{
		"id", "realm_id", "event_id", "type", "inserted_at",
	}

	return &pgEventStore{
		adapter: db.NewCRUDTable(session, "", EventTableName, defaultColumns, sortableColumns),
	}
}

// -----------------------------------------------------------------------------

func (r *pgEventStore) Create(ctx context.Context, entity *models.Event) error {
	return r.adapter.Create(ctx, entity)
}
