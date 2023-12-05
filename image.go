package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

// readImage reads an image file and returns it
func readImage(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	return image.Decode(f)
}

// writeImage writes an image to a file
func writeImage(filename string, img *image.Gray) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

// imageToSlice returns the selected pixels of the image in a slice
func imageToSlice(img *image.Gray) []uint8 {
	s := []uint8{}
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.GrayAt(x, y)
			s = append(s, c.Y)
		}
	}

	return s
}

// sliceImage returns the pixels of an image region as a slice of grayscale values
func sliceImage(filename string, bounds image.Rectangle, writeImg bool) ([]byte, error) {
	img, _, err := readImage(filename)
	if err != nil {
		return nil, err
	}

	// Draw a grayscale image of the portion of interest
	grayscale := image.NewGray(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: bounds.Dx(), Y: bounds.Dy()},
	})
	draw.Draw(grayscale, grayscale.Bounds(), img, bounds.Min, draw.Src)

	if writeImg {
		err = writeImage("out_"+filename, grayscale)
		if err != nil {
			return nil, err
		}
	}

	return imageToSlice(grayscale), nil
}
