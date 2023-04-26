package connection

import (
	"context"
	_net "notifier/pkg/net"
	"notifier/pkg/ulid"
	"time"
)

func NewInteractor(repo Repository) Usecase {
	return &interactor{
		repo:   repo,
		expire: 30 * time.Minute,
	}
}

type interactor struct {
	repo   Repository
	expire time.Duration
}

func (i interactor) CreateConnection(c context.Context, id ulid.ID) (*Connection, error) {
	e := NewConnection(id, _net.GetAddress(), i.expire)
	if err := i.repo.Save(c, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (i interactor) CloseConnection(c context.Context, id ulid.ID) error {
	if err := i.repo.Delete(c, id); err != nil {
		return err
	}
	return nil
}

func (i interactor) GetConnection(c context.Context, id ulid.ID) (*Connection, error) {
	e, err := i.repo.FetchOne(c, id)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (i interactor) ExtendConnection(c context.Context, id ulid.ID) error {
	if err := i.repo.ExtendTTL(c, id, i.expire); err != nil {
		return err
	}
	return nil
}

func (i interactor) GenerateID() ulid.ID {
	return ulid.NewID()
}
