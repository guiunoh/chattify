package publish

import (
	"context"
	"fmt"
	"log"
	"publisher/pkg/ulid"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

func NewRepository(cache *redis.Client, channel string) Repository {
	return &repository{
		cache:   cache,
		channel: channel,
	}
}

type repository struct {
	cache   *redis.Client
	channel string
}

func (r repository) Key(id ulid.ID) string {
	return fmt.Sprintf("%s:%s", r.channel, id.String())
}

func (r repository) Save(ctx context.Context, item *Publish) error {
	value, err := json.Marshal(item)
	if err != nil {
		return err
	}
	result, err := r.cache.Publish(ctx, r.channel, value).Result()
	if err != nil {
		return err
	}

	log.Println("publish:", result)
	return nil
}
