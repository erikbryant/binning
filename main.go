package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"image"
	"log"
)

var (
	filename = flag.String("filename", "", "Filename to bin")
	x        = flag.Int("x", 277, "Upper left corner")
	y        = flag.Int("y", 100, "Upper left corner")
	width    = flag.Int("width", 1000, "Width")
	height   = flag.Int("height", 60, "Height")
	debug    = flag.Bool("debug", false, "Enable debug output of image section")
	echo     = flag.Bool("echo", false, "Echo filename to output")
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("  Bin an image by a computing a checksum of [a portion of] the image")
	fmt.Println("    binning -filename 'file.png' -x 0 -y 10 -width 100 -height 100")
}

func coordsToBounds(minX, minY int, width, height int) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{Y: minY, X: minX},
		Max: image.Point{Y: minY + height, X: minX + width},
	}
}

func main() {
	flag.Parse()

	if *filename == "" {
		usage()
		log.Fatal("You must supply a filename")
	}

	bounds := coordsToBounds(*x, *y, *width, *height)

	slice, err := sliceImage(*filename, bounds, *debug)
	if err != nil {
		log.Fatal(err)
	}

	crc := crc32.ChecksumIEEE(slice)

	fmt.Print(crc)
	if *echo {
		fmt.Print(" ", *filename)
	}
	fmt.Println()
}
