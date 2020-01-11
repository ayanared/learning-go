package main

import "fmt"

func main() {
	const x string = "The girl " // declare + initialize
	fmt.Println(x)
	y := "with the" // same as above - just a shortcut
	fmt.Println(y)
	var a = "yellow hat is happy" // initialize, declaration is implicit
	fmt.Println(a)
	var b, c int = 1, 2 //declare + initialize multiple variables
	fmt.Println(b, c)
}
