package main

import (
	"image/color"
	"fmt"
	"strconv"
)

type Rgba color.RGBA

func (clr Rgba) ToString() string {
	return fmt.Sprintf("rgb(%s, %s, %s)", strconv.Itoa(int(clr.R)), strconv.Itoa(int(clr.G)), strconv.Itoa(int(clr.B)))
}

func (clr Rgba) ToRgba() color.RGBA {
	return color.RGBA{R: clr.R, G: clr.G, B: clr.B, A: clr.A}
}

var Colours map[string]string = map[string]string{
	"rgb(232, 25, 19)": "A",
	"rgb(250, 124, 12)": "B",
	"rgb(242, 176, 5)": "C",
	"rgb(247, 227, 54)": "D",
	"rgb(155, 200, 54)": "E",
	"rgb(67, 175, 238)": "F",
	"rgb(30, 59, 171)": "G",
	"rgb(101, 40, 104)": "H",
	"rgb(229, 50, 108)": " I",
	"rgb(251, 175, 217)": "J",
}