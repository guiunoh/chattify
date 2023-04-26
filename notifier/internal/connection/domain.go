package connection

import (
	"context"
	"notifier/pkg/ulid"
	"time"
)

type Repository interface {
	Key(id ulid.ID) string
	Save(c context.Context, connection *Connection) error
	Delete(c context.Context, id ulid.ID) error
	FetchOne(c context.Context, id ulid.ID) (*Connection, error)
	ExtendTTL(c context.Context, id ulid.ID, expiration time.Duration) error
}

type Usecase interface {
	CreateConnection(c context.Context, id ulid.ID) (*Connection, error)
	CloseConnection(c context.Context, id ulid.ID) error
	GetConnection(c context.Context, id ulid.ID) (*Connection, error)
	ExtendConnection(c context.Context, id ulid.ID) error
	GenerateID() ulid.ID
}
