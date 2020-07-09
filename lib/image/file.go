package image

import (
	"image"
	"io/ioutil"
	"optimus/lib/webp"
)

// File represents an image file.
type File struct {
	Data     []byte `json:"data"`
	Ext      string `json:"ext"`
	MimeType string `json:"type"`
	Name     string `json:"name"`
	Image    image.Image
}

func (f *File) Write() error {
	buf, err := webp.EncodeWebp(f.Image)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(f.Name+".webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
