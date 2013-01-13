package main

// Still in progress

import (
  "tour/tree"
	"fmt"
	)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	switch {
	case t.Left != nil:
		Walk(t.Left, ch)
		ch <- t.Value
		if (t.Right != nil){
			Walk(t.Right, ch)
		}
	default:
		ch <- t.Value
		if (t.Right != nil){
			Walk(t.Right, ch)
		}
		
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	var x, y int
	
	select {
	case c1 <- x:
		c2 <- y
		if x != y {
			return false 
		}
	case c2<- y:
		c1 <- x
		if x != y {
			return false
		}
	default:
		return true
	}
	return true
}

func main() {
/*
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
*/
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
