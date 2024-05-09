package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++{
		line := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			line[j] = uint8(i ^ j)
		}
		p[i] = line
	}
	return p
}

func main() {
	pic.Show(Pic)
}
