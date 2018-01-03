package pic2ascii

import (
	"image"
	"image/color"
)

type Resize struct {
	image.Image
	x, y   int
	dx, dy int
}

// Resize image
func NewResize(img image.Image, x, y int) image.Image {
	if x == 0 && y == 0 {
		return img
	}
	rect := img.Bounds()
	dx := rect.Dx()
	dy := rect.Dy()
	if x == 0 {
		x = y * dx / dy
	} else if y == 0 {
		y = x * dy / dx
	}
	return Resize{img, x, y, dx, dy}
}

func (c Resize) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{c.x, c.y},
	}
}

func (c Resize) At(x, y int) color.Color {
	ats := []color.Color{}
	j0, l0 := y*c.dy/c.y, (y+1)*c.dy/c.y
	i0, k0 := x*c.dx/c.x, (x+1)*c.dx/c.x
	for i, k := i0, k0; i != k; i++ {
		for j, l := j0, l0; j != l; j++ {
			ats = append(ats, c.Image.At(i, j))
		}
	}
	if len(ats) == 0 {
		return c.Image.At(i0, j0)
	}
	return Sum(ats)
}

func Sum(cs []color.Color) color.Color {
	switch len(cs) {
	case 0:
		return color.White
	case 1:
		return cs[0]
	}
	var sr, sg, sb, sa uint32 = 0, 0, 0, 0
	for _, v := range cs {
		r, g, b, a := v.RGBA()
		sr += r
		sg += g
		sb += b
		sa += a
	}
	cl := uint32(len(cs))
	sr /= cl
	sg /= cl
	sb /= cl
	sa /= cl
	return color.RGBA64{uint16(sr), uint16(sg), uint16(sb), uint16(sa)}
}
