package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/nfnt/resize"
	"github.com/wzshiming/pic2ascii"
	"gopkg.in/ffmt.v1"
)

func main() {
	pic := flag.String("p", "", "input image file")
	chars := flag.String("c", `MMNXKAVQOCL#kxdoclc\=;:"'.  `, "chars")
	w := flag.Uint("w", 0, "resize width")
	h := flag.Uint("h", 0, "resize height")
	o := flag.String("o", "", "output file")
	flag.Parse()

	if *pic == "" {
		flag.Usage()
		return
	}

	u, err := url.Parse(*pic)
	if err != nil {
		ffmt.Mark(err)
		return
	}

	var f []byte

	switch u.Scheme {
	case "http", "https":
		resp, err := http.Get(u.String())
		if err != nil {
			ffmt.Mark(err)
			return
		}
		defer resp.Body.Close()

		f, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			ffmt.Mark(err)
			return
		}
	case "file", "":
		f, err = ioutil.ReadFile(u.Path)
		if err != nil {
			ffmt.Mark(err)
			return
		}
	default:
		ffmt.Mark("unknown scheme ", u.Scheme)
		return
	}

	img, _, err := image.Decode(bytes.NewBuffer(f))
	if err != nil {
		ffmt.Mark(err)
		return
	}

	if *w != 0 || *h != 0 {
		img = resize.Resize(*w, *h, img, resize.Lanczos3)
	}

	dd := string(pic2ascii.ToAscii(img, []rune(*chars)))

	if *o == "" {
		fmt.Print(dd)
	} else {
		ioutil.WriteFile(*o, []byte(dd), 0666)
	}

}
