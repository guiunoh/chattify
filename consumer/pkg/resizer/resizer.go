package resizer

import "image"

type ImageResizer interface {
	Resize(src image.Image, rate float32) ([]byte, string, error)
}
