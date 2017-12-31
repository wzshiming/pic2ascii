package pic2ascii

import (
	"bytes"
	"image"
)

//图片转为字符画
func ToAscii(m image.Image, arr []rune) string {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	t := (256 * 256) / len(arr)
	dst := bytes.NewBuffer(nil)
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			cr := m.At(j, i)
			r, g, b, _ := cr.RGBA()
			avg := float64(r) + float64(g) + float64(b)
			avg /= 3.0
			num := int(avg) / t
			dst.WriteRune(arr[num])
		}
		dst.WriteRune('\n')
	}
	return dst.String()
}
