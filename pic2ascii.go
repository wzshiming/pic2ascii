package pic2ascii

import (
	"image"
)

// Image to Ascii
func ToAscii(m image.Image, arr []rune) []rune {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	t := (256 * 256) / (len(arr) - 1)
	dst := make([]rune, 0, (dy * (dx + 1)))
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			cr := m.At(j, i)
			r, g, b, _ := cr.RGBA()
			avg := float64(r) + float64(g) + float64(b)
			avg /= 3.0
			num := int(avg) / t
			dst = append(dst, arr[num])
		}
		dst = append(dst, '\n')
	}
	return dst
}
