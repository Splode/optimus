package webp

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"optimus/lib/image"

	"github.com/chai2010/webp"
)

func Write(i *image.File) error {
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
	if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}

func Toast() {
	var buf bytes.Buffer
	// var width, height int
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./CNY_Artwork-en.jpg"); err != nil {
		log.Println(err)
	}

	m, err := jpeg.Decode(bytes.NewReader(data))

	// GetInfo
	// if width, height, _, err = webp.GetInfo(data); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("width = %d, height = %d\n", width, height)

	// GetMetadata
	// if metadata, err := webp.GetMetadata(data, "ICCP"); err != nil {
	// 	fmt.Printf("Metadata: err = %v\n", err)
	// } else {
	// 	fmt.Printf("Metadata: %s\n", string(metadata))
	// }

	// Decode webp
	// m, err := webp.Decode(bytes.NewReader(data))
	// if err != nil {
	// 	log.Println(err)
	// }
	// Decode webp
	// m, err := webp.DecodeRGBA(bytes.NewReader(data))
	// if err != nil {
	// 	log.Println(err)
	// }

	// m, err := webp.Load("./test.jpg")
	// if err != nil {
	// 	log.Println(err)
	// }
	// Encode lossless webp
	if err = webp.Encode(&buf, m, &webp.Options{Lossless: false, Quality: 70}); err != nil {
		log.Println(err)
	}
	if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	fmt.Println("Save output.webp ok")
}
