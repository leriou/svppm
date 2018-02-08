# 图片保存

将特定格式的数据写入图片

该项目模仿 milo yip的svpng构建

svpng: https://github.com/miloyip/svpng

使用方法:

```go
func main() {
	row := make([]int, 256*256*3)
	p := 0
	for i := 0; i <= 255; i++ {
		for j := 0; j <= 255; j++ {
			row[p] = j
			row[p+1] = i
			row[p+2] = 128
			p += 3
		}
	}
	svppm.Svppm("test.ppm", 256, 256, row)
	fmt.Println("well,done")
}

```