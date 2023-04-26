package message

import (
	"publisher/pkg/ulid"
	"time"
)

type Message struct {
	ID       ulid.ID   `json:"id" gorm:"primaryKey"`
	Topic    string    `json:"topic"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	ExpiryAt time.Time `json:"expiryAt" gorm:"column:expiry_at"`
	CreateAt time.Time `json:"createAt" gorm:"column:create_at"`
}

func NewMessage(topic, author, content string, expire time.Duration) *Message {
	createAt := time.Now()
	return &Message{
		ID:       ulid.NewID(),
		Topic:    topic,
		Author:   author,
		Content:  content,
		CreateAt: createAt,
		ExpiryAt: createAt.Add(expire),
	}
}

// TableName Tabler interface
func (Message) TableName() string {
	return "message"
}
