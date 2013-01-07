package main

/*
Implement WordCount. It should return a map of the counts of each .word. in the string s.
http://tour.golang.org/#45
*/

import (
	"fmt"
	"strings"
	"os"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	for _,word := range strings.Fields(s) {
		m[word]++
	}

	return m
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(WordCount("hello hello world"))
	} else {
		fmt.Println(WordCount(strings.Join(os.Args[1:], " ")))
	}

	
}

