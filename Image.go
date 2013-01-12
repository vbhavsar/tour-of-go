package main

// Remember the picture generator you wrote earlier? Let's write another one, 
// but this time it will return an implementation of image.Image instead of a slice of data.
// http://tour.golang.org/#58


import (
  "image"
	"tour/pic"
	"image/color"
)

type Image struct{}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
