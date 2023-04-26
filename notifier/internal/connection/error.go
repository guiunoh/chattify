package connection

import "github.com/pkg/errors"

var (
	ErrDuplicateKey = errors.New("notifier: duplicate key")
	ErrNotFound     = errors.New("notifier: not found")
)
