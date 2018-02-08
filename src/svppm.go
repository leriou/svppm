package svppm

import (
	"fmt"
	"os"
	// "strings"
	"strconv"
)

const (
	MAXDEEP = 255
)

func Svppm(f string, w, h int, row []int) bool {
	fi, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	header := "P3\n" + strconv.Itoa(w) + " " + strconv.Itoa(h) + "\n" + strconv.Itoa(MAXDEEP) + "\n"
	var rowdata string
	for p := 0; p < len(row); p += 3 {
		rowdata += strconv.Itoa(row[p]) + " " + strconv.Itoa(row[p+1]) + " " + strconv.Itoa(row[p+2]) + "\n"
	}
	content := header + rowdata
	fi.WriteString(content)
	fmt.Println("write success")
	return true
}

