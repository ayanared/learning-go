package main

import "fmt"

//A slice, unlike an array is dynamically-sized,
func main() {

	x := make([]string, 3) //create an empty slice with non-zero length
	fmt.Println(x)

	x[0] = "x"
	x[1] = "y"
	x[2] = "z"
	fmt.Println("length:", len(x))
	//add to end of slice
	x = append(x, "a")
	x = append(x, "b")
	x = append(x, "c")
	fmt.Println("length:", len(x))
	fmt.Println(x)

	primes := [6]int{2, 3, 5, 7, 11, 13} // array
	var t []int = primes[1:4]            // slice

	fmt.Printf("%T\n", t)

}
