package models

import (
	"time"

	"github.com/dchest/uniuri"
)

// Event describes event attribubes for database
type Event struct {
	RealmID   string    `json:"realm_id" db:"realm_id"`
	EventID   string    `json:"event_id" db:"event_id"`
	Type      string    `json:"type" db:"type"`
	Payload   []byte    `json:"payload" db:"payload"`
	Timestamp time.Time `json:"timestamp" db:"inserted_at"`
}

// NewEvent returns a initialized event instance
func NewEvent() *Event {
	return &Event{
		EventID:   uniuri.NewLen(64),
		Timestamp: time.Now().UTC(),
	}
}

// -----------------------------------------------------------------------------
