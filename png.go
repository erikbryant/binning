package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// readImage reads an image file and returns it as a grayscale image
func readImage(filename string) (*image.Gray, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// Convert the image to grayscale
	grayscale := image.NewGray(img.Bounds())
	draw.Draw(grayscale, grayscale.Bounds(), img, image.Point{}, draw.Src)

	return grayscale, nil
}

// imageToSlice returns the selected pixels of the image in a slice
func imageToSlice(img *image.Gray, bounds image.Rectangle) []uint8 {
	s := []uint8{}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.GrayAt(x, y)
			s = append(s, c.Y)
		}
	}

	return s
}

func writeSlice(filename string, src []byte, bounds image.Rectangle) error {
	// Create a blank image to receive the pixels
	img := image.NewGray(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	// Copy pixels from the source region
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			i := y*bounds.Dx() + x
			img.Set(x, y, color.Gray{
				Y: src[i],
			})
		}
	}

	// Write new image to a file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}

	return nil
}

func slicePNG(filename string, bounds image.Rectangle, writeImg bool) ([]byte, error) {
	img, err := readImage(filename)
	if err != nil {
		return nil, err
	}

	slice := imageToSlice(img, bounds)

	if writeImg {
		err = writeSlice("out_"+filename, slice, bounds)
		if err != nil {
			return nil, err
		}
	}

	return slice, nil
}
