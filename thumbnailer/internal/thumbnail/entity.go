package thumbnail

import (
	"fmt"
	"thumbnailer/pkg/ulid"
	"time"
)

type Thumbnail struct {
	ID     ulid.ID   `json:"id"`
	Type   string    `json:"type"`
	Data   []byte    `json:"data"`
	Expiry time.Time `json:"expiry"`
}

func NewThumbnail(data []byte, imageType string) *Thumbnail {
	const expiry = 3 * time.Minute
	return &Thumbnail{
		ID:     ulid.NewID(),
		Data:   data,
		Type:   imageType,
		Expiry: time.Now().Add(expiry),
	}
}

func (e Thumbnail) Key() string {
	return fmt.Sprintf("%s:%s", e.TableName(), e.ID.String())
}

// TableName Tabler interface
func (e Thumbnail) TableName() string {
	return "thumbnail"
}
