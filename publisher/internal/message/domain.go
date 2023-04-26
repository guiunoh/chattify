package message

import (
	"context"
	"publisher/pkg/ulid"
)

type Repository interface {
	Save(ctx context.Context, message *Message) error
	FetchOne(ctx context.Context, id ulid.ID, topic string) (*Message, error)
	Fetch(ctx context.Context, lastID ulid.ID, topic string, limit int, forward bool) ([]Message, error)
}

type Writer interface {
	CreateMessage(c context.Context, topic string, author, content string) (*Message, error)
}
type Usecase interface {
	Writer
	GetMessageOne(c context.Context, id ulid.ID, topic string) (*Message, error)
	GetMessageMany(c context.Context, id ulid.ID, topic string, limit int, forward bool) ([]Message, error)
}
