package main

// Calculates sqrt using Newton's method
// http://tour.golang.org/#44

import (
	"fmt"
	"math"
)

const delta = 0.0005

func Sqrt(x float64) float64 {
	if x == 0 {
		return x
	}
	sqrt_act := math.Sqrt(x)

	z := x/2
	
	diff := math.Abs(sqrt_act - z)
	runs := 0
	
	for diff > delta {
		z = z - ((z*z - x) / 2*z)
		diff = math.Abs(sqrt_act - z)
		runs++
	}
	fmt.Println("runs",runs)
	return z;
}

func main() {
	sqrt_cal := Sqrt(2)
	sqrt_act := math.Sqrt(2)
	
	fmt.Println("calculated:",sqrt_cal)
	fmt.Println("actual:",sqrt_act)
	fmt.Println("diff:",sqrt_cal - sqrt_act);
}
