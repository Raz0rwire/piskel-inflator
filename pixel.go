package main

import (
	"io"
	"image"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func RgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)}
}

func ReadPixels(file io.Reader) ([][]Pixel, int, int, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := make([][]Pixel, 0, height * EnlargementFactor)
	for y := 0; y < height; y++ {
		row := make([]Pixel, 0, width * EnlargementFactor)
		for x := 0; x < width; x++ {
			row = append(row, RgbaToPixel(img.At(x, y).RGBA()))
		}

		pixels = append(pixels, row)
	}

	return pixels, width, height, nil
}
