//+build support_video

package pic2ascii

import (
	"errors"
	"image"
	"io"
	"time"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/cgo/ffmpeg"
	"github.com/nareix/joy4/format"
)

var rc io.Reader
var iotoken = "mem://video"

// 兼容 github.com/nareix/joy4 一个强制断言的 bug

type noper struct {
	r io.Reader
	s []byte
	i int64
}

func (r *noper) sycn(i int) (int64, error) {
	return io.CopyN(r, r.r, int64(i))
}

func (r *noper) Write(p []byte) (n int, err error) {
	r.s = append(r.s, p...)
	return len(p), nil
}

func (r *noper) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		if i, _ := r.sycn(1024 * 10); i != 0 {
			return r.Read(b)
		}
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *noper) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.i + offset
		//	case io.SeekEnd:
		//		abs = int64(len(r.s)) + offset
	default:
		return 0, errors.New("pic2ascii.noper.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("pic2ascii.noper.Seek: negative position")
	}
	r.i = abs
	return abs, nil
}

func (r *noper) Close() error { return nil }

func newNoper(r io.Reader) *noper { return &noper{r, nil, 0} }

func init() {
	avutil.DefaultHandlers.Add(func(r *avutil.RegisterHandler) {
		r.UrlReader = func(f string) (bool, io.ReadCloser, error) {
			if f != iotoken {
				return false, nil, nil
			}

			if rc0 := rc; rc0 != nil {
				rc = nil
				return true, newNoper(rc0), nil
			}
			return false, nil, nil
		}
	})
	format.RegisterAll()
}

func VideoSlice(read io.Reader, f func(time.Duration, image.Image)) error {
	rc = read
	file, err := avutil.Open(iotoken)
	if err != nil {
		return err
	}
	defer file.Close()

	streams, err := file.Streams()
	if err != nil {
		return err
	}

	var vd *ffmpeg.VideoDecoder
	var inx int
	for ind, s := range streams {
		if s.Type().IsVideo() {
			vd0, err := ffmpeg.NewVideoDecoder(s)
			if err != nil {
				return err
			}
			vd = vd0
			inx = ind
			break
		}
	}

	for {
		pkt, err := file.ReadPacket()
		if err != nil {
			if io.EOF == err {
				break
			}
			return err
		}

		if pkt.Idx != int8(inx) {
			continue
		}

		vf, err := vd.Decode(pkt.Data)
		if err != nil {
			return err
		}

		f(pkt.Time, &vf.Image)
	}
	return nil
}
