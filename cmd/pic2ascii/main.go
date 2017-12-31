package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"github.com/nfnt/resize"
	"github.com/wzshiming/pic2ascii"
	"gopkg.in/ffmt.v1"
)

func main() {
	pic := flag.String("pic", "", "image file")
	chars := flag.String("chars", `MNXKAVQOCL#kxdoclc\=;:"'. `, "chars")
	w := flag.Uint("width", 0, "resize width")
	h := flag.Uint("height", 0, "resize height")
	o := flag.String("out", "", "out file")
	flag.Parse()

	f, err := ioutil.ReadFile(*pic)
	if err != nil {
		ffmt.Mark(err)
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

	dd := pic2ascii.ToAscii(img, []rune(*chars))

	if *o == "" {
		fmt.Print(dd)
	} else {
		ioutil.WriteFile(*o, []byte(dd), 0666)
	}

}
