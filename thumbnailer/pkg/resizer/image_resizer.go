package resizer

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
)

type imageResizer struct{}

func NewImageResizer() ImageResizer {
	return &imageResizer{}
}

func (r *imageResizer) Resize(src image.Image, rate float32) ([]byte, string, error) {
	width := int(float32(src.Bounds().Dx()) * rate)
	dst := imaging.Resize(src, width, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dst); err != nil {
		panic(err)
	}

	return buf.Bytes(), "png", nil
}
