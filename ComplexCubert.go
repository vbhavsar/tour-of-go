package main

import (
  "fmt"
	"math"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	
	z := x
	
	for i:=0; i<10; i++ {
		z = z - ( (cmplx.Pow(z,3)-x) / (3*cmplx.Pow(z,2)))
	}
	
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(math.Cbrt(2))
}
