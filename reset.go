package pic2ascii

import (
	"image"
	"image/color"
)

type Reset struct {
	image.Image
	ox, oy int
	dx, dy int
}

// Reset image
func NewReset(img image.Image) image.Image {
	rect := img.Bounds()

	ox := rect.Min.X
	oy := rect.Min.Y

	if ox == 0 && oy == 0 {
		return img
	}

	dx := rect.Dx()
	dy := rect.Dy()

	return Reset{img, ox, oy, dx, dy}
}

func (c Reset) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{c.dx, c.dy},
	}
}

func (c Reset) At(x, y int) color.Color {
	return c.Image.At(x+c.ox, y+c.oy)
}
