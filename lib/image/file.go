package image

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
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
		f.Image, err = jpeg.Decode(bytes.NewReader(f.Data))
	case "png":
		f.Image, err = png.Decode(bytes.NewReader(f.Data))
	case "webp":
		f.Image, err = webp.DecodeWebp(bytes.NewReader(f.Data))
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

// Write saves a file to disk based on the encoding target.
func (f *File) Write(dir string, target string) (err error) {
	var buf bytes.Buffer
	switch target {
	case "jpg":
		err = jpeg.Encode(&buf, f.Image, &jpeg.Options{Quality: 70})
	case "png":
		err = png.Encode(&buf, f.Image)
	case "webp":
		buf, err = webp.EncodeWebp(f.Image)
	}
	dest := path.Join(dir, f.Name+"."+target)
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
