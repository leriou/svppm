package main

import (
	"fmt"
	svppm "svppm/src"
)

const (
	WIDTH  = 256
	HEIGHT = 256
)

func main() {
	row := make([]int, WIDTH*HEIGHT*3)
	p := 0
	r := 0
	g := 0
	b := 0
	for i := 0; i < WIDTH; i++ {
		r += 1
		for j := 0; j < HEIGHT; j++ {
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
	svppm.Svppm("test.ppm", WIDTH, HEIGHT, row)
	fmt.Println("well,done")
}
