package jpeg

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
)

// Options represent JPEG encoding options.
type Options struct {
	Quality int `json:"quality"`
}

// DecodeJPEG decodes a JPEG file and return an image.
func DecodeJPEG(r io.Reader) (image.Image, error) {
	i, err := jpeg.Decode(r)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// EncodeJPEG encodes an image into JPEG and returns a buffer.
func EncodeJPEG(i image.Image, o *Options) (buf bytes.Buffer, err error) {
	err = jpeg.Encode(&buf, i, &jpeg.Options{Quality: o.Quality})
	return buf, err
}
