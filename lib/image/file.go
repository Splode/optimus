package image

import (
	"image"
	"io/ioutil"
	"optimus/lib/webp"
	"path"
)

// File represents an image file.
type File struct {
	Data     []byte `json:"data"`
	Ext      string `json:"ext"`
	MimeType string `json:"type"`
	Name     string `json:"name"`
	Image    image.Image
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
	return nil
}
