package subscribe

import (
	"publisher/pkg/ulid"
	"time"
)

type Subscribe struct {
	ID           ulid.ID   `json:"id" gorm:"primaryKey;"`
	ConnectionID ulid.ID   `json:"connectionID" gorm:"column:connection_id"`
	Topic        string    `json:"topic"`
	ExpiryAt     time.Time `json:"expiryAt" gorm:"column:expiry_at"`
	CreateAt     time.Time `json:"createAt" gorm:"column:create_at"`
}

func NewSubscribe(connectionID ulid.ID, topic string) *Subscribe {
	const expiry = 24 * time.Hour
	createAt := time.Now()
	return &Subscribe{
		ID:           ulid.NewID(),
		ConnectionID: connectionID,
		Topic:        topic,
		CreateAt:     createAt,
		ExpiryAt:     createAt.Add(expiry),
	}
}

// TableName Tabler interface
func (Subscribe) TableName() string {
	return "subscribe"
}
