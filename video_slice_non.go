//+build !support_video

package pic2ascii

import (
	"image"
	"io"
	"time"
)

var VideoSlice func(read io.Reader, f func(time.Duration, image.Image)) error
