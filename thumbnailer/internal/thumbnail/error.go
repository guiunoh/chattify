package thumbnail

import "github.com/pkg/errors"

var (
	ErrThumbnailNotFound     = errors.New("thumbnailer: thumbnail not found")
	ErrThumbnailDuplicateKey = errors.New("thumbnailer: duplicate key")
)
