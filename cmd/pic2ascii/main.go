package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

	buf := bytes.NewReader(f)
	switch *t {
	case "gif":
		err = showGIF(buf)
	default:
		err = showElse(buf)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	return
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
	if *o != "" {
		return ioutil.WriteFile(*o, []byte(dd), 0666)
	}

	fmt.Print(dd)
	return nil
}

func toAscii(img image.Image) string {
	img = pic2ascii.NewReset(img)

	if *w != 0 || *h != 0 {
		img = pic2ascii.NewResize(img, int(*w), int(*h))
	}

	return string(pic2ascii.ToAscii(img, []rune(*chars), []rune(*prefix), []rune(*suffix)))
}

func getFile(filpath string) ([]byte, error) {
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

		f, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return f, nil
	case "file", "":
		f, err := ioutil.ReadFile(u.Path)
		if err != nil {
			return nil, err
		}
		return f, nil
	default:
		return nil, fmt.Errorf("unknown scheme %v", u.Scheme)
	}
}
