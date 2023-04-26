package subscribe

import (
	"context"
	"publisher/pkg/ulid"
)

func NewInteractor(r Repository) Usecase {
	return &interactor{r}
}

type interactor struct {
	repo Repository
}

func (i interactor) GetSubscribesBytTopic(c context.Context, topic string) ([]Subscribe, error) {
	items, err := i.repo.FetchByTopic(c, topic)
	if err != nil {
		return nil, err
	}

	return items, err
}

func (i interactor) Subscribe(c context.Context, connectionID ulid.ID, topic string) error {
	e := NewSubscribe(connectionID, topic)
	if err := i.repo.Save(c, e); err != nil {
		return err
	}
	return nil
}

func (i interactor) UnSubscribe(c context.Context, connectionID ulid.ID, topic string) error {
	if len(topic) > 0 {
		e, err := i.repo.FetchOne(c, connectionID, topic)
		if err != nil {
			return err
		}

		if err := i.repo.DeleteOne(c, e); err != nil {
			return err
		}
		return nil
	}

	if err := i.repo.DeleteByConnectionID(c, connectionID); err != nil {
		return err
	}
	return nil
}

func (i interactor) UnSubscribeAllByConnectionID(c context.Context, connectionID ulid.ID) {
	//TODO implement me
	panic("implement me")
}
