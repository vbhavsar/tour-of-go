package main

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers.
// http://tour.golang.org/#47

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  prev1, prev0 := 0, 1
	
	return func() int { 
		sum := prev1+prev0
		prev1 = prev0 
		prev0 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
