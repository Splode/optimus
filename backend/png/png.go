package png

import (
	"bytes"
	"image"
	"image/png"
	"io"
)

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
// Lossless compression is used; the Quality option controls the compression
// level (higher quality = less compression effort = faster encode).
func EncodePNG(i image.Image, o *Options) (buf bytes.Buffer, err error) {
	enc := &png.Encoder{CompressionLevel: compressionLevel(o.Quality)}
	err = enc.Encode(&buf, i)
	return buf, err
}

// compressionLevel maps a quality value (0–100) to a png.CompressionLevel.
// Lower quality = more aggressive compression (smaller file, slower encode).
func compressionLevel(q int) png.CompressionLevel {
	switch {
	case q >= 80:
		return png.BestSpeed
	case q >= 40:
		return png.DefaultCompression
	default:
		return png.BestCompression
	}
}
