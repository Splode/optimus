package image

import (
	"bytes"
	"errors"
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
	var buf bytes.Buffer
	switch c.App.Target {
	case "jpg":
		buf, err = jpeg.EncodeJPEG(f.Image, c.App.JpegOpt)
	case "png":
		buf, err = png.EncodePNG(f.Image, c.App.PngOpt)
	case "webp":
		buf, err = webp.EncodeWebp(f.Image, c.App.WebpOpt)
	}
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

// getFileType returns the file's type based on the given mime type.
func getFileType(t string) (string, error) {
	m, prs := mimes[t]
	if !prs {
		_ = errors.New("unsupported file type:" + t)
	}
	return m, nil
}
