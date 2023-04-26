package message

import (
	"context"
	"publisher/pkg/ulid"
	"time"
)

func NewInteractor(r Repository) Usecase {
	return &interactor{
		repo:   r,
		expire: (365 * 24) * time.Hour,
	}
}

type interactor struct {
	repo   Repository
	expire time.Duration
}

func (i interactor) CreateMessage(c context.Context, topic, author, content string) (*Message, error) {
	message := NewMessage(topic, author, content, i.expire)

	if err := i.repo.Save(c, message); err != nil {
		return nil, err
	}

	return message, nil
}

func (i interactor) GetMessageOne(c context.Context, id ulid.ID, topic string) (*Message, error) {
	message, err := i.repo.FetchOne(c, id, topic)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (i interactor) GetMessageMany(c context.Context, id ulid.ID, topic string, limit int, forward bool) ([]Message, error) {
	messages, err := i.repo.Fetch(c, id, topic, limit, forward)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
