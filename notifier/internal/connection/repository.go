package connection

import (
	"context"
	"fmt"
	"notifier/pkg/ulid"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

func NewRepository(cache *redis.Client, prefix string) Repository {
	return &repository{
		cache:  cache,
		prefix: prefix,
	}
}

type repository struct {
	cache  *redis.Client
	prefix string
}

func (r repository) Key(id ulid.ID) string {
	return fmt.Sprintf("%s:%s", r.prefix, id.String())
}

func (r repository) Save(c context.Context, connection *Connection) error {
	key := r.Key(connection.ID)
	value, err := json.Marshal(connection)
	if err != nil {
		return err
	}
	expirationTime := connection.ExpiryAt.Sub(connection.CreateAt)

	result, err := r.cache.SetNX(c, key, value, expirationTime).Result()
	if err != nil {
		return err
	}

	if !result {
		return ErrDuplicateKey
	}

	return nil
}

func (r repository) Delete(c context.Context, id ulid.ID) error {
	key := r.Key(id)
	_, err := r.cache.Del(c, key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r repository) FetchOne(c context.Context, id ulid.ID) (*Connection, error) {
	key := r.Key(id)
	value, err := r.cache.Get(c, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound

		}
		return nil, err
	}

	var e Connection
	err = json.Unmarshal(value, &e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r repository) ExtendTTL(c context.Context, id ulid.ID, expiration time.Duration) error {
	key := r.Key(id)
	if err := r.cache.Expire(c, key, expiration).Err(); err != nil {
		return err
	}
	return nil
}
