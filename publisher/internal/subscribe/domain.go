package subscribe

import (
	"context"
	"publisher/pkg/ulid"
)

type Repository interface {
	Save(ctx context.Context, subscribe *Subscribe) error
	FetchOne(ctx context.Context, connectionID ulid.ID, topic string) (*Subscribe, error)
	FetchByTopic(ctx context.Context, topic string) ([]Subscribe, error)
	DeleteOne(ctx context.Context, subscribe *Subscribe) error
	DeleteByConnectionID(ctx context.Context, connectionID ulid.ID) error
}

type Reader interface {
	GetSubscribesBytTopic(c context.Context, topic string) ([]Subscribe, error)
}

type Usecase interface {
	Reader
	Subscribe(c context.Context, connectionID ulid.ID, topic string) error
	UnSubscribe(c context.Context, connectionID ulid.ID, topic string) error
	UnSubscribeAllByConnectionID(c context.Context, connectionID ulid.ID)
}
