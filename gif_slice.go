package pic2ascii

import (
	"image"
	"image/draw"
	"image/gif"
)

// SliceGIF
func SliceGIF(g *gif.GIF) (r []image.Image) {
	for _, v := range g.Image {
		if len(r) == 0 {
			r = append(r, v)
		} else {
			r = append(r, MergeImage(v, r[len(r)-1]))
		}
	}
	return
}

// MergeImage
func MergeImage(src image.Image, over image.Image) image.Image {
	dist := image.NewNRGBA(over.Bounds())
	draw.Draw(dist, dist.Bounds(), over, image.ZP, draw.Src)
	draw.Draw(dist, dist.Bounds(), src, image.ZP, draw.Over)
	return dist
}
