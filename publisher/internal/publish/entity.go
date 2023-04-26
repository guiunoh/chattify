package publish

import "publisher/pkg/ulid"

type Publish struct {
	ConnectionID ulid.ID `json:"connectionID"`
	Topic        string  `json:"topic"`
	Payload      string  `json:"payload"`
}

func NewPublish(connectionID ulid.ID, topic string, payload string) *Publish {
	return &Publish{connectionID, topic, payload}
}
