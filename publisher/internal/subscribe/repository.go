package subscribe

import (
	"context"
	"publisher/pkg/ulid"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Save(ctx context.Context, subscribe *Subscribe) error {
	if err := r.db.WithContext(ctx).Create(subscribe).Error; err != nil {
		return err
	}
	return nil
}

func (r repository) FetchOne(ctx context.Context, connectionID ulid.ID, topic string) (*Subscribe, error) {
	var e Subscribe
	if err := r.db.WithContext(ctx).
		Where("connection_id = ? AND topic = ?", connectionID, topic).
		Find(&e).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func (r repository) FetchByTopic(ctx context.Context, topic string) ([]Subscribe, error) {
	var items []Subscribe
	if err := r.db.WithContext(ctx).
		Where("topic = ?", topic).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r repository) DeleteOne(ctx context.Context, subscribe *Subscribe) error {
	if err := r.db.WithContext(ctx).Delete(subscribe).Error; err != nil {
		return err
	}
	return nil
}

func (r repository) DeleteByConnectionID(ctx context.Context, connectionID ulid.ID) error {
	if err := r.db.WithContext(ctx).
		Where("connection_id = ?", connectionID).
		Delete(&Subscribe{}).Error; err != nil {
		return err
	}
	return nil
}
