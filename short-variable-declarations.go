package main

import "fmt"

var i = 1

// This is not available.
// j := 1

func main() {
	var x, y int = 1, 2
	k := 3
	c, java, python := true, false, "No!"
	fmt.Println(i, x, y, k, c, java, python)
}
