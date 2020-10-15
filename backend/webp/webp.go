package webp

import (
	"bytes"
	"github.com/chai2010/webp"
	"image"
	"io"
)

// Options represent WebP encoding options.
type Options struct {
	Lossless bool `json:"lossless"`
	Quality  int  `json:"quality"`
}

// DecodeWebp a webp file and return an image.
func DecodeWebp(r io.Reader) (image.Image, error) {
	i, err := webp.Decode(r)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// EncodeWebp encodes an image into webp and returns a buffer.
func EncodeWebp(i image.Image, o *Options) (buf bytes.Buffer, err error) {
	if err = webp.Encode(&buf, i, &webp.Options{Lossless: o.Lossless, Quality: float32(o.Quality)}); err != nil {
		return buf, err
	}
	return buf, nil
}
