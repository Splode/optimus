package image

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/wailsapp/wails"
	"image"
	"io/ioutil"
	"optimus/lib/config"
	"optimus/lib/jpeg"
	"optimus/lib/png"
	"optimus/lib/webp"
	"os"
	"path"
	"path/filepath"
)

const (
	FILL = iota
	FIT
)

var mimes = map[string]string{
	"image/.jpg": "jpg",
	"image/jpg":  "jpg",
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/webp": "webp",
}

// File represents an image file.
type File struct {
	Data          []byte `json:"data"`
	Ext           string `json:"ext"`
	ID            string `json:"id"`
	MimeType      string `json:"type"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	ConvertedFile string
	IsConverted   bool
	Image         image.Image
	Runtime       *wails.Runtime
}

// Decode decodes the file's data based on its mime type.
func (f *File) Decode() error {
	mime, err := getFileType(f.MimeType)
	if err != nil {
		return err
	}

	switch mime {
	case "jpg":
		f.Image, err = jpeg.DecodeJPEG(bytes.NewReader(f.Data))
	case "png":
		f.Image, err = png.DecodePNG(bytes.NewReader(f.Data))
	case "webp":
		f.Image, err = webp.DecodeWebp(bytes.NewReader(f.Data))
	}
	if err != nil {
		return err
	}
	return nil
}

// GetConvertedSize returns the size of the converted file.
func (f *File) GetConvertedSize() (int64, error) {
	if f.ConvertedFile == "" {
		return 0, errors.New("file has no converted file")
	}
	s, err := os.Stat(f.ConvertedFile)
	if err != nil {
		return 0, err
	}
	return s.Size(), nil
}

// GetSavings returns the delta between original and converted file size.
func (f *File) GetSavings() (int64, error) {
	c, err := f.GetConvertedSize()
	if err != nil {
		return 0, err
	}
	return f.Size - c, nil
}

// Write saves a file to disk based on the encoding target.
func (f *File) Write(c *config.Config) (err error) {
	// TODO resizing should probably be in its own method
	if c.App.Sizes != nil {
		for _, r := range c.App.Sizes {
			if r.Height <= 0 || r.Width <= 0 {
				f.Runtime.Events.Emit("notify", map[string]interface{}{
					"msg":  fmt.Sprintf("Invalid image size: %s", r.String()),
					"type": "warn",
				})
				continue
			}
			var i image.Image
			switch r.Strategy {
			case FILL:
				i = imaging.Fill(f.Image, r.Width, r.Height, imaging.Center, imaging.Lanczos)
			case FIT:
				i = imaging.Fit(f.Image, r.Width, r.Height, imaging.Lanczos)
				// TODO determine file name size based on returned image rect
			}
			buf, err := encToBuf(i, c.App)
			dest := path.Join(c.App.OutDir, c.App.Prefix+f.Name+"--"+r.String()+c.App.Suffix+"."+c.App.Target)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(dest, buf.Bytes(), 0666); err != nil {
				return err
			}
		}
	}
	buf, err := encToBuf(f.Image, c.App)
	dest := path.Join(c.App.OutDir, c.App.Prefix+f.Name+c.App.Suffix+"."+c.App.Target)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dest, buf.Bytes(), 0666); err != nil {
		return err
	}
	f.ConvertedFile = filepath.Clean(dest)
	f.IsConverted = true
	return nil
}

// encToBuf encodes an image to a buffer using the configured target.
func encToBuf(i image.Image, a *config.App) (*bytes.Buffer, error) {
	var b bytes.Buffer
	var err error
	switch a.Target {
	case "jpg":
		b, err = jpeg.EncodeJPEG(i, a.JpegOpt)
	case "png":
		b, err = png.EncodePNG(i, a.PngOpt)
	case "webp":
		b, err = webp.EncodeWebp(i, a.WebpOpt)
	}
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// getFileType returns the file's type based on the given mime type.
func getFileType(t string) (string, error) {
	m, prs := mimes[t]
	if !prs {
		_ = errors.New("unsupported file type:" + t)
	}
	return m, nil
}
