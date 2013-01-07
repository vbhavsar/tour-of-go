package main

/*
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
http://tour.golang.org/#46
*/

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	rv := make([][]uint8, dx)
	for x:=0; x<dx; x++{
		rv[x] = make([]uint8, dy) 
		for y:=0; y<dy; y++{
			rv[x][y] = uint8(x^y);
		}
	}
	return rv;
}

func main() {
	pic.Show(Pic)
}
