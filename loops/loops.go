package main

import (
	"fmt"
)

//constants can't be changed.
//A numeric constant has no type until it is given one, such as by explict conversion
//So you can think of it as un-typed...

func main() {
	//single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classic initial condition after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//without any conditions (keeps going until break)
	for {
		fmt.Println("loop")
		break
	}

	//using the keyword "continue"
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
