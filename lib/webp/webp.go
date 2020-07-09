package webp

import (
	"bytes"
	"github.com/chai2010/webp"
	"image"
)

// EncodeWebp encodes an image into webp and returns a buffer.
func EncodeWebp(i image.Image) (buf bytes.Buffer, err error) {
	if err = webp.Encode(&buf, i, &webp.Options{Lossless: false, Quality: 70}); err != nil {
		return buf, err
	}
	return buf, nil
}
