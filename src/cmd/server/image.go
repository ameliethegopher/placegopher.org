package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"

	"github.com/muesli/smartcrop"
	"github.com/nfnt/resize"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func resizeImage(name string, width, height int, colour string) error {
	f, err := os.Open("img/" + name + ".jpg")
	if err != nil {
		return err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	analyzer := smartcrop.NewAnalyzer()
	topCrop, err := analyzer.FindBestCrop(img, width, height)
	if err != nil {
		return err
	}

	// The crop will have the requested aspect ratio, but you need to copy/scale it yourself
	fmt.Printf("Top crop: %+v\n", topCrop)

	// get the part of the image, as specified by topCrop
	cropped := img.(SubImager).SubImage(topCrop)

	// resize this 'cropped' part to the dimensions we require
	resized := resize.Resize(uint(width), uint(height), cropped, resize.Bicubic)

	// write out to disk
	writeImage("jpeg", resized, "out/"+name+"-"+strconv.Itoa(width)+"x"+strconv.Itoa(height)+".jpg")

	return nil
}

// Functions from : https://github.com/muesli/smartcrop/blob/master/debug.go

func writeImage(imgtype string, img image.Image, name string) error {
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		panic(err)
	}

	switch imgtype {
	case "png":
		return writeImageToPng(img, name)
	case "jpeg":
		return writeImageToJpeg(img, name)
	}

	return errors.New("Unknown image type")
}

func writeImageToJpeg(img image.Image, name string) error {
	fso, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fso.Close()

	return jpeg.Encode(fso, img, &jpeg.Options{Quality: 100})
}

func writeImageToPng(img image.Image, name string) error {
	fso, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fso.Close()

	return png.Encode(fso, img)
}
