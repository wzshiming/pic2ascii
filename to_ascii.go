package pic2ascii

import (
	"image"
	"image/color"
)

const (
	max = 1<<16 - 1
)

// Image to Ascii
func ToAscii(m image.Image, chars []rune, prefix, suffix []rune) []rune {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	t := max / (len(chars) - 1)
	dst := make([]rune, 0, (dy * (dx + len(prefix) + len(suffix))))
	for i := 0; i != dy; i++ {
		if len(prefix) != 0 {
			dst = append(dst, prefix...)
		}
		for j := 0; j != dx; j++ {
			cr := m.At(j, i)
			g, _, _, _ := color.GrayModel.Convert(color.NRGBAModel.Convert(cr)).RGBA()
			ii := int(g) / t
			dst = append(dst, chars[ii])
		}
		if len(suffix) != 0 {
			dst = append(dst, suffix...)
		}
	}
	return dst
}
