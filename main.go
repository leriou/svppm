package main

import (
	"fmt"
	svppm "svppm/src"
)

func main() {
	row := make([]int, 256*256*3)
	p := 0
	r := 0
	g := 0
	b := 0
	for i := 0; i <= 255; i++ {
		r += 1
		for j := 0; j <= 255; j++ {
			g += 1
			row[p] = r
			row[p+1] = g
			if r > g {
				b = r
			} else {
				b = g
			}
			row[p+2] = b
			p += 3
		}
	}
	svppm.Svppm("test.ppm", 256, 256, row)
	fmt.Println("well,done")
}
