package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/wzshiming/pic2ascii"

	_ "image/gif"
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

	img, _, err := image.Decode(bytes.NewBuffer(f))
	if err != nil {
		fmt.Println(err)
		return
	}

	if *w != 0 || *h != 0 {
		img = pic2ascii.NewResize(img, int(*w), int(*h))
	}

	if *r {
		*chars = reverseString(*chars)
	}
	dd := string(pic2ascii.ToAscii(img, []rune(*chars), []rune(*prefix), []rune(*suffix)))

	if *o == "" {
		fmt.Print(dd)
	} else {
		ioutil.WriteFile(*o, []byte(dd), 0666)
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
