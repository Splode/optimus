package image

import (
	"bytes"
	"encoding/json"
	"github.com/chai2010/webp"
	"image"
	"image/jpeg"
	"io/ioutil"
)

// File represents an image file.
type File struct {
	Data     []byte `json:"data"`
	MimeType string `json:"type"`
	Name     string `json:"name"`
	Image    image.Image
}

// HandleFile processes a file submission from the client.
func HandleFile(fileJson string) (file File, err error) {
	if err := json.Unmarshal([]byte(fileJson), &file); err != nil {
		return file, err
	}
	file.Image, err = jpeg.Decode(bytes.NewReader(file.Data))
	if err != nil {
		return file, nil
	}
	if err := Write(&file); err != nil {
		return file, err
	}
	return file, nil
}

func Write(i *File) error {
	var buf bytes.Buffer
	//var data []byte
	var err error

	// Load file data
	//if data, err = ioutil.ReadFile("./test.jpg"); err != nil {
	//	log.Println(err)
	//}
	//
	//m, err := jpeg.Decode(bytes.NewReader(data))

	// Encode lossless webp
	if err = webp.Encode(&buf, i.Image, &webp.Options{Lossless: false, Quality: 70}); err != nil {
		return err
	}
	if err = ioutil.WriteFile("test.webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
