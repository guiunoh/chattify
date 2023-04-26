package connection

import (
	"net"
	"notifier/pkg/ulid"
	"time"
)

type Connection struct {
	ID       ulid.ID   `json:"id"`
	Host     net.IP    `json:"host"`
	ExpiryAt time.Time `json:"expiryAt"`
	CreateAt time.Time `json:"createAt"`
}

func NewConnection(id ulid.ID, host net.IP, expire time.Duration) *Connection {
	createAt := time.Now()
	return &Connection{
		ID:       id,
		Host:     host,
		ExpiryAt: createAt.Add(expire),
		CreateAt: createAt,
	}
}
