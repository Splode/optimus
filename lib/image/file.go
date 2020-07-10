package image

import (
	"errors"
	"image"
	"io/ioutil"
	"optimus/lib/webp"
	"os"
	"path"
)

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

func (f *File) Write(dir string) error {
	buf, err := webp.EncodeWebp(f.Image)
	dest := path.Join(dir, f.Name+".webp")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dest, buf.Bytes(), 0666); err != nil {
		return err
	}
	f.ConvertedFile = dest
	return nil
}
