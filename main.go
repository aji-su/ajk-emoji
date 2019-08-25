package main

import (
	"fmt"
	"image"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("original.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	src = imaging.CropAnchor(src, 320, 480, imaging.Bottom)
	err = imaging.Save(src, "out.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	var arrs [][]int

	for y := 0; y < 480; y += 32 {
		var yarrs []int
		for x := 0; x < 320; x += 32 {
			yarrs = append(yarrs, x)
		}
		arrs = append(arrs, yarrs)
	}

	for i, ys := range arrs {
		for j, x := range ys {
			dst := imaging.Crop(src, image.Rect(x, i*32, x+32, i*32+32))
			ename := fmt.Sprintf("zzajk%02d%d", i, j)
			err = imaging.Save(dst, ename+".png")
			if err != nil {
				log.Fatalf("failed to save image: %v", err)
			}
			fmt.Printf(":%s:\u200b", ename)
			if j == 9 {
				fmt.Println()
			}
		}
	}
}
