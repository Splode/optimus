package png

import (
	"bytes"
	"github.com/foobaz/lossypng/lossypng"
	"image"
	"image/png"
	"io"
)

// TODO: pass options to compressor

// Decode a PNG file and return an image.
func DecodePNG(r io.Reader) (image.Image, error) {
	i, err := png.Decode(r)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// EncodePNG encodes an image into PNG and returns a buffer.
func EncodePNG(i image.Image) (buf bytes.Buffer, err error) {
	c := lossypng.Compress(i, -1, 10)
	err = png.Encode(&buf, c)
	return buf, err
}
