package publish

import (
	"context"
	"publisher/pkg/ulid"
)

type Repository interface {
	Key(id ulid.ID) string
	Save(ctx context.Context, item *Publish) error
}

type Usecase interface {
	PostToTopic(c context.Context, topic, sender, content string) error
}
