package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
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

func main() {
	pic := flag.String("f", "", "input image file")
	chars := flag.String("c", `MMWNXK0Okxou=:"'.  `, "chars")
	r := flag.Bool("r", false, "reverse chars")
	w := flag.Uint("w", 0, "resize width")
	h := flag.Uint("h", 0, "resize height")
	o := flag.String("o", "", "output file")
	t := flag.String("t", "", "file type")
	m := flag.Int("m", -1, "Gif max loop count")
	prefix := flag.String("p", "", "prefix")
	suffix := flag.String("s", "\n", "suffix")
	flag.Parse()

	if *pic == "" {
		flag.Usage()
		return
	}

	u, err := url.Parse(*pic)
	if err != nil {
		fmt.Println(err)
		return
	}

	var f []byte

	switch u.Scheme {
	case "http", "https":
		resp, err := http.Get(u.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		f, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "file", "":
		f, err = ioutil.ReadFile(u.Path)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("unknown scheme ", u.Scheme)
		return
	}

	if *r {
		*chars = reverseString(*chars)
	}

	toAscii := func(img image.Image) string {
		img = pic2ascii.NewReset(img)

		if *w != 0 || *h != 0 {
			img = pic2ascii.NewResize(img, int(*w), int(*h))
		}

		return string(pic2ascii.ToAscii(img, []rune(*chars), []rune(*prefix), []rune(*suffix)))
	}

	buf := bytes.NewReader(f)

	if *t == "" {
		*t = strings.TrimPrefix(filepath.Ext(*pic), ".")
	}
	*t = strings.ToLower(*t)

	switch *t {
	case "gif":
		img, err := gif.DecodeAll(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		dds := []string{}
		sg := pic2ascii.SliceGIF(img)
		for _, v := range sg {
			dds = append(dds, fmt.Sprintln(toAscii(v)))
		}

		if *o == "" {
			if img.LoopCount == 0 {
				img.LoopCount = *m
			}

			for i := 0; i != img.LoopCount; i++ {
				for k, v := range dds {
					fmt.Println(v)
					time.Sleep(time.Duration(img.Delay[k]) * time.Second / 100)
				}
			}
		} else {
			ioutil.WriteFile(*o, []byte(strings.Join(dds, "\n")), 0666)
		}

	default:
		img, _, err := image.Decode(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		dd := toAscii(img)
		if *o == "" {
			fmt.Print(dd)
		} else {
			ioutil.WriteFile(*o, []byte(dd), 0666)
		}
	}
}

func reverseString(s string) string {
	str := []rune(s)
	l := len(str) / 2
	for i := 0; i < l; i++ {
		j := len(str) - i - 1
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
