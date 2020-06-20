package main

import (
	"fmt"
	"math"
	"os"
)

const (
	MAXDEEP = 255
	WIDTH   = 256
	HEIGHT  = 256
)

func Svppm(filename string, w, h int, row []int) bool {
	c := fmt.Sprintf("P3\n%d %d\n%d\n", w, h, MAXDEEP)
	for p := 0; p < len(row); p += 3 {
		c += fmt.Sprintf("%d %d %d\n", row[p], row[p+1], row[p+2])
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(c)
	fmt.Println("write success !!!")
	return true
}

func main() {
	row := make([]int, WIDTH*HEIGHT*3)
	r := 0
	g := 0
	p := 0
	for i := 0; i < WIDTH; i++ {
		r += 1
		for j := 0; j < HEIGHT; j++ {
			g += 1
			row[p] = r
			row[p+1] = g
			row[p+2] = int(math.Min(float64(r), float64(g)))
			p += 3
		}
	}
	Svppm("test.ppm", WIDTH, HEIGHT, row)
	fmt.Println("well,done")
}
