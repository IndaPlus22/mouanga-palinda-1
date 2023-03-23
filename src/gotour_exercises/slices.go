package main

import "golang.org/x/tour/pic"

// probably mixed up my dys and dxes
func Pic(dx, dy int) [][]uint8 {
	list := make([][]uint8, dy)
	for i := 0; i < dx; i += 1 {
		list[i] = make([]uint8, dx)
	}

	for i := 0; i < dx; i += 1 {
		for j := 0; j < dy; j += 1 {
			list[i][j] = uint8((i + j) / 2)
		}
	}
	return list
}

func main() {

	pic.Show(Pic)

}
