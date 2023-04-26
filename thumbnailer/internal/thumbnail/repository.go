package thumbnail

import (
	"context"
	"fmt"
	"thumbnailer/pkg/ulid"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

func NewRepository(rdb *redis.Client) Repository {
	return &repository{rdb, "thumbnail"}
}

type repository struct {
	rdb    *redis.Client
	prefix string
}

func (r *repository) key(id ulid.ID) string {
	return fmt.Sprintf("%s:%s", r.prefix, id.String())
}

func (r *repository) Save(ctx context.Context, entity *Thumbnail) error {
	key := r.key(entity.ID)
	value, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	expirationTime := time.Until(entity.Expiry)

	result, err := r.rdb.SetNX(ctx, key, value, expirationTime).Result()
	if err != nil {
		return err
	}

	if !result {
		return ErrThumbnailDuplicateKey
	}

	return nil
}

func (r *repository) FindByID(ctx context.Context, id ulid.ID) (*Thumbnail, error) {
	key := r.key(id)
	value, err := r.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrThumbnailNotFound

		}
		return nil, err
	}

	var entity Thumbnail
	err = json.Unmarshal(value, &entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
