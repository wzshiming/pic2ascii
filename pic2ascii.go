package pic2ascii

import (
	"image"
	"image/color"
)

var (
	Suffix = []rune("\n")
)

const (
	max = 1<<16 - 1
)

// Image to Ascii
func ToAscii(m image.Image, chars []rune) []rune {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	t := max / (len(chars) - 1)
	dst := make([]rune, 0, (dy * (dx + len(Suffix))))
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			cr := m.At(j, i)
			g, _, _, _ := color.NRGBAModel.Convert(color.GrayModel.Convert(cr)).RGBA()
			ii := int(g) / t
			dst = append(dst, chars[ii])
		}
		dst = append(dst, Suffix...)
	}
	return dst
}
