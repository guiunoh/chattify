package connection

import (
	"fmt"
	"notifier/pkg/ulid"
	"time"
)

type Presenter interface {
	Connected(id ulid.ID, expiry time.Time) any
	Usage() any
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p presenter) Connected(id ulid.ID, expiry time.Time) any {
	maxAge := int(time.Until(expiry).Seconds())
	return map[string]any{
		"connectionID": id,
		"cacheControl": fmt.Sprintf("maxAge=%d", maxAge),
	}
}

func (p presenter) Usage() any {
	return map[string]any{
		"forwarder":   "{\"type\" : \"forwarder\", \"topic\": \"[topic-id]\"}",
		"unSubscribe": "{\"type\" : \"unsubscribe\", \"topic\": \"[topic-id]\"}",
	}
}
