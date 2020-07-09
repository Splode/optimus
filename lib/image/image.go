package image

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"optimus/lib/webp"
)

// File represents an image file.
type File struct {
	Data     []byte `json:"data"`
	MimeType string `json:"type"`
	Name     string `json:"name"`
	Image    image.Image
}

// HandleFile processes a file from the client.
func HandleFile(fileJson string) (file File, err error) {
	if err := json.Unmarshal([]byte(fileJson), &file); err != nil {
		return file, err
	}

	if file.MimeType == "image/jpg" {
		file.Image, err = jpeg.Decode(bytes.NewReader(file.Data))
	} else if file.MimeType == "image/png" {
		file.Image, err = png.Decode(bytes.NewReader(file.Data))
	}
	if err != nil {
		return file, nil
	}

	if err := file.Write(); err != nil {
		return file, err
	}
	return file, nil
}

func (f *File) Write() error {
	buf, err := webp.EncodeWebp(f.Image)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("test.webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
