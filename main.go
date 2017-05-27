package main

import (
	"fmt"
	"image"
	"os"
	"image/color"
	"github.com/fogleman/gg"
	"github.com/disintegration/imaging"
)

const EnlargementFactor = 33
const EveryNthGuideIsThickened = 5

func main() {
	file, err := os.Open("image.png")
	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, width, height, err := ReadPixels(file)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	EnlargeAndLabelImage(pixels, width, height)
}

func EnlargeAndLabelImage(pixels [][]Pixel, width int, height int) string {
	max_x := width * EnlargementFactor
	max_y := height * EnlargementFactor

	mask := image.NewRGBA(image.Rect(0, 0, max_x, max_y))
	lineMask := image.NewRGBA(image.Rect(0, 0, max_x, max_y))
	textMask := gg.NewContext(max_x, max_y)

	/*
	Uncomment and load your desired font if you want

	err := textMask.LoadFontFace("/Users/kevinhoogerwerf/Library/Fonts/Roboto-Black.ttf", 18)
	if (err != nil ) {
		panic(err)
	}*/

	for row := 0; row < len(pixels); row++ {
		thickenedRow := row % EveryNthGuideIsThickened == EveryNthGuideIsThickened - 1

		columns := pixels[row];
		for column := 0; column < len(columns); column++ {
			thickenedColumn := column % EveryNthGuideIsThickened == EveryNthGuideIsThickened - 1

			clr := columns[column];

			x := column * EnlargementFactor;
			y := row * EnlargementFactor

			newColor := Rgba{R: clr.R, G: clr.G, B: clr.B, A: clr.A}
			for i_y := y; i_y < (y + EnlargementFactor); i_y++ {

				for i_x := x; i_x < (x + EnlargementFactor); i_x++ {

					if ( i_y + 1 == max_y || i_y + 2 == max_y) {
						//Last horizontal fat line
						lineMask.Set(i_x, i_y, color.RGBA{0, 0, 0, 255})
					} else if ( i_x + 1 == max_x || i_x + 2 == max_x ) {
						//Last vertical fat line
						lineMask.Set(i_x, i_y, color.RGBA{0, 0, 0, 255})
					} else if (i_x + 1 == x + EnlargementFactor || thickenedColumn && i_x + 2 == x + EnlargementFactor || thickenedColumn && i_x + 3 == x + EnlargementFactor) {
						lineMask.Set(i_x, i_y, color.RGBA{0, 0, 0, 255})
					} else if (i_y + 1 == y + EnlargementFactor || (thickenedRow && i_y + 2 == y + EnlargementFactor) || (thickenedRow && i_y + 3 == y + EnlargementFactor)) {
						lineMask.Set(i_x, i_y, color.RGBA{0, 0, 0, 255})
					} else if ( i_y == 0) {
						//First horizontal fat line
						lineMask.Set(i_x, i_y, color.RGBA{0, 0, 0, 255})
					} else if ( i_x == 0) {
						//First vertical fat line
						lineMask.Set(i_x, i_y, color.RGBA{0, 0, 0, 255})
					}

					mask.Set(i_x, i_y, newColor.ToRgba())
				}
			}

			letter := Colours[newColor.ToString()]

			textMask.SetRGB(1, 1, 1)
			textMask.DrawString(letter, float64(x - 6 + (EnlargementFactor / 2)), float64(y + 7 + (EnlargementFactor / 2)))

		}
	}

	n := imaging.FlipH(lineMask.SubImage(image.Rect(0, 0, max_x, max_y)))

	final := gg.NewContext(max_x, max_y)
	final.DrawImage(mask, 0, 0)
	final.DrawImage(n, 0, 0)
	final.DrawImage(textMask.Image(), 0, 0)

	final.SavePNG("final.png")
	return "final.png";

}