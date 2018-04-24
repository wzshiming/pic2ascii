package pic2ascii

import (
	"image"
	"image/draw"
	"image/gif"
)

// SliceGIF
func SliceGIF(g *gif.GIF, f func(image.Image)) {
	if len(g.Image) == 0 {
		return
	}
	var img image.Image = g.Image[0]
	f(img)
	if len(g.Image) == 1 {
		return
	}

	for _, v := range g.Image[1:] {
		img = MergeImage(img, v)
		f(img)
	}
	return
}

// MergeImage
func MergeImage(src image.Image, over image.Image) image.Image {
	dist := image.NewRGBA(src.Bounds())
	draw.Draw(dist, dist.Bounds(), src, image.ZP, draw.Src)
	draw.Draw(dist, dist.Bounds(), over, image.ZP, draw.Over)
	return dist
}
