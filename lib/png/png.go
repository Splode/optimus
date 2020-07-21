package png

import (
	"bytes"
	"github.com/foobaz/lossypng/lossypng"
	"image"
	"image/png"
	"io"
)

// TODO: pass options to compressor

// Options represent PNG encoding options.
type Options struct {
	Quality int `json:"quality"`
}

// DecodePNG decodes a PNG file and return an image.
func DecodePNG(r io.Reader) (image.Image, error) {
	i, err := png.Decode(r)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// EncodePNG encodes an image into PNG and returns a buffer.
func EncodePNG(i image.Image, o *Options) (buf bytes.Buffer, err error) {
	c := lossypng.Compress(i, -1, o.Quality)
	err = png.Encode(&buf, c)
	return buf, err
}
