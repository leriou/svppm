package svppm

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MAXDEEP = 255
)

func Svppm(filename string, w, h int, row []int) bool {
	c := "P3\n" + strconv.Itoa(w) + " " + strconv.Itoa(h) + "\n" + strconv.Itoa(MAXDEEP) + "\n"
	for p := 0; p < len(row); p += 3 {
		c += strconv.Itoa(row[p]) + " " + strconv.Itoa(row[p+1]) + " " + strconv.Itoa(row[p+2]) + "\n"
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
