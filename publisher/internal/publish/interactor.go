package publish

import (
	"context"
	"log"
	"publisher/internal/message"
	"publisher/internal/subscribe"
	"publisher/pkg/ulid"
)

func NewInteractor(repo Repository, reader subscribe.Reader, writer message.Writer) Usecase {
	return &interactor{repo, reader, writer}
}

type interactor struct {
	repo   Repository
	reader subscribe.Reader
	writer message.Writer
}

func (i interactor) PostToTopic(c context.Context, topic, sender, content string) error {
	m, err := i.writer.CreateMessage(c, topic, sender, content)
	log.Println("message:", m)

	subscribes, err := i.reader.GetSubscribesBytTopic(c, topic)
	if err != nil {
		return err
	}

	for _, s := range subscribes {
		go func(connectionID ulid.ID, topic string, payload string) {
			e := NewPublish(connectionID, topic, payload)
			log.Println("publish:", e)
			if err := i.repo.Save(context.Background(), e); err != nil {
				log.Println("error:", err)
				return
			}
		}(s.ConnectionID, topic, content)
	}

	return nil
}
