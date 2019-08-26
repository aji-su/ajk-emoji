package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"path"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

func main() {
	var outprefix string
	flag.StringVar(&outprefix, "op", "img/zz_", "output filepath prefix")

	var outsuffix string
	flag.StringVar(&outsuffix, "os", ".png", "output filepath suffix")

	var xsplit int
	flag.IntVar(&xsplit, "x", 10, "number of columns")

	var outputHTML bool
	flag.BoolVar(&outputHTML, "html", false, "outputs html")

	flag.Parse()

	log.Print()

	ret := 0
	if err := run(outprefix, outsuffix, xsplit, outputHTML); err != nil {
		ret = 1
	}

	os.Exit(ret)
}

func run(outprefix, outsuffix string, xsplit int, outputHTML bool) error {
	src, err := imaging.Decode(os.Stdin)
	if err != nil {
		return err
	}

	rct := src.Bounds()

	psize := rct.Dx() / xsplit
	ysplit := rct.Dy() / psize
	log.Printf("psize=%d,xsplit=%d,ysplit=%d", psize, xsplit, ysplit)

	var emojis, html string

	for i := 0; i < ysplit; i++ {
		y := i * psize
		for j := 0; j < xsplit; j++ {
			x := j * psize
			dst := imaging.Crop(src, image.Rect(x, y, x+psize, y+psize))
			ename := fmt.Sprintf(outprefix+"%02d%02d", i, j)
			if err := imaging.Save(dst, ename+outsuffix); err != nil {
				return err
			}
			emojis += fmt.Sprintf(":%s:\u200b", path.Base(ename))
			cwd, err := os.Getwd()
			if err != nil {
				return nil
			}
			html += fmt.Sprintf(`<img src="%s/%s%s" />`, cwd, ename, outsuffix)
			if j == xsplit-1 {
				emojis += fmt.Sprintf("\n")
				html += "<br>"
			}
		}
	}

	if outputHTML {
		fmt.Printf("<style>img{border:solid 1px black;width:48px}</style>\n%s\n<pre>%s</pre>\n", html, emojis)
	} else {
		fmt.Println(emojis)
	}
	return nil
}
