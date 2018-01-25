package pic2ascii

import (
	"image"
	"image/draw"
	"image/gif"
)

// SliceGIF
func SliceGIF(g *gif.GIF) (r []image.Image) {
	if len(g.Image) == 0 {
		return nil
	}
	r = make([]image.Image, 0, len(g.Image))
	r = append(r, g.Image[0])
	if len(g.Image) == 1 {
		return r
	}

	for _, v := range g.Image[1:] {
		r = append(r, MergeImage(r[len(r)-1], v))
	}
	return r
}

// MergeImage
func MergeImage(src image.Image, over image.Image) image.Image {
	dist := image.NewRGBA(src.Bounds())
	draw.Draw(dist, dist.Bounds(), src, image.ZP, draw.Src)
	draw.Draw(dist, dist.Bounds(), over, image.ZP, draw.Over)
	return dist
}
