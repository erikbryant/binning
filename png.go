package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

// readPNG reads a PNG file and returns it as a grayscale image
func readPNG(filename string) (*image.Gray, error) {
	// Open the PNG file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Decode the RGB PNG file into an image
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// Convert the RGB image to grayscale
	grayscale := image.NewGray(img.Bounds())
	draw.Draw(grayscale, grayscale.Bounds(), img, image.Point{}, draw.Src)

	return grayscale, nil
}

func writePNG(filename string, src *image.Gray, bounds image.Rectangle) error {
	// Create a blank image to receive the pixels
	img := image.NewGray(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	// Copy pixels from the source region
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			srcX := x + bounds.Min.X
			srcY := y + bounds.Min.Y
			img.Set(x, y, src.At(srcX, srcY))
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

func slicePNG(filename string, bounds image.Rectangle) ([]byte, error) {
	img, err := readPNG(filename)
	if err != nil {
		return nil, err
	}

	err = writePNG("out_"+filename, img, bounds)
	if err != nil {
		return nil, err
	}

	return imageToSlice(img, bounds), nil
}
