package main

import (
	"fmt"
	"math"
)

//constants can't be changed.
//A numeric constant has no type until it is given one, such as by explict conversion
//So you can think of it as un-typed...

func main() {
	const n = 50000000
	fmt.Printf("%T\n", n) //int
	//math.Sin(n) expects a float64 so n is a float64 below
	fmt.Println(math.Sin(n)) //0.8256467432733234
	fmt.Printf("%T\n", n)    //int
}
