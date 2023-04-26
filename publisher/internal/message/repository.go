package message

import (
	"context"
	"publisher/pkg/ulid"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Save(ctx context.Context, message *Message) error {
	if err := r.db.WithContext(ctx).Create(&message).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) FetchOne(ctx context.Context, id ulid.ID, topic string) (*Message, error) {
	var message Message
	if err := r.db.WithContext(ctx).
		Where("id = ? AND topic = ?", id, topic).
		Find(&message).Error; err != nil {
		return nil, err
	}

	return &message, nil
}

func (r repository) Fetch(ctx context.Context, lastID ulid.ID, topic string, limit int, forward bool) ([]Message, error) {
	var messages []Message
	condition, order := r.idConditionAndOrder(forward)
	if err := r.db.WithContext(ctx).
		Where(condition, lastID).
		Where("topic = ?", topic).
		Order(order).
		Limit(limit).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r repository) idConditionAndOrder(forward bool) (condition string, order string) {
	const (
		idGreater = "id > ?"
		idLess    = "id < ?"
		idDesc    = "id DESC"
		idAsc     = "id ASC"
	)

	if forward {
		return idLess, idAsc
	}
	return idGreater, idDesc
}
