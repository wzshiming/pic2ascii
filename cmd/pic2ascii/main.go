package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/wzshiming/pic2ascii"

	"image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var (
	pic    = flag.String("f", "", "input image file")
	chars  = flag.String("c", `MMWNXK0Okxou=:"'.  `, "chars")
	r      = flag.Bool("r", false, "reverse chars")
	w      = flag.Uint("w", 0, "resize width")
	h      = flag.Uint("h", 0, "resize height")
	o      = flag.String("o", "", "output file")
	t      = flag.String("t", "", "file type")
	m      = flag.Int("m", -1, "Gif max loop count")
	prefix = flag.String("p", "", "prefix")
	suffix = flag.String("s", "\n", "suffix")
)

func init() {
	flag.Parse()
	if *r {
		*chars = pic2ascii.ReverseString(*chars)
	}
	if *t == "" {
		*t = strings.TrimPrefix(filepath.Ext(*pic), ".")
	}
	*t = strings.ToLower(*t)
}

func main() {

	if *pic == "" {
		flag.Usage()
		return
	}

	f, err := getFile(*pic)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	switch *t {
	case "gif":
		err = showGIF(f)
	case "mp4", "ts", "rtmp", "rtsp", "flv", "aac":
		if pic2ascii.VideoSlice != nil {
			err = showVideo(f)
		} else {
			err = fmt.Errorf("The current version does not support video.")
		}
	default:
		//	case "jpeg", "jpg", "png", "bmp", "tiff", "webp":
		err = showElse(f)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func showVideo(buf io.ReadCloser) error {
	var sum time.Duration
	dds := []string{}
	err := pic2ascii.VideoSlice(buf, func(dur time.Duration, img image.Image) {
		v := toAscii(img)
		fmt.Println(v)
		time.Sleep(dur - sum)
		sum = dur
		dds = append(dds, v)
	})
	if err != nil {
		return err
	}
	if *o != "" {
		return ioutil.WriteFile(*o, []byte(strings.Join(dds, "\n")), 0666)
	}
	return nil
}

func showGIF(buf io.Reader) error {
	img, err := gif.DecodeAll(buf)
	if err != nil {
		return err
	}

	dds := []string{}
	pic2ascii.SliceGIF(img, func(v image.Image) {
		dd := toAscii(v)
		fmt.Println(dd)
		time.Sleep(time.Duration(img.Delay[len(dds)]) * time.Second / 100)
		dds = append(dds, dd)
	})

	if *o != "" {
		return ioutil.WriteFile(*o, []byte(strings.Join(dds, "\n")), 0666)
	}

	if *m != 0 {
		img.LoopCount = *m
	}

	for i := 0; i != img.LoopCount; i++ {
		for k, v := range dds {
			fmt.Println(v)
			time.Sleep(time.Duration(img.Delay[k]) * time.Second / 100)
		}
	}
	return nil
}

func showElse(buf io.Reader) error {
	img, _, err := image.Decode(buf)
	if err != nil {
		return err
	}

	dd := toAscii(img)
	fmt.Print(dd)

	if *o != "" {
		return ioutil.WriteFile(*o, []byte(dd), 0666)
	}
	return nil
}

func toAscii(img image.Image) string {
	img = pic2ascii.NewReset(img)

	if *w != 0 || *h != 0 {
		img = pic2ascii.NewResize(img, int(*w), int(*h))
	}

	return string(pic2ascii.ToAscii(img, []rune(*chars), []rune(*prefix), []rune(*suffix)))
}

func getFile(filpath string) (io.ReadCloser, error) {
	u, err := url.Parse(filpath)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		cli := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		resp, err := cli.Get(u.String())
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	case "file", "":
		file, err := os.OpenFile(u.Path, os.O_RDONLY, 0)
		if err != nil {
			return nil, err
		}
		return file, nil
	default:
		return nil, fmt.Errorf("unknown scheme %v", u.Scheme)
	}
}
