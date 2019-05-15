package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red color.Color = color.RGBA{200, 30, 30, 255}

func main() {
	file, err := os.Create("someimage.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)
	// или draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)

	for x := 1; x <= 4; x++ {
		for y := 20; y <= 480; y++ {
			img.Set(x*100, y, red)
			img.Set(y, x*100, red)
		}
	}

	png.Encode(file, img)
}
