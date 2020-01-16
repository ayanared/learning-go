package main

import (
	"fmt"
	"strconv"
)

//func name-of-function(input-variable type) return-type{
//	do something with the input-variable
//	return whatever
//}
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func namedReturnValues(val1, val2 int) (ret1, ret2 string) {
	ret1 = ("The first input is " + strconv.Itoa(val1))
	ret2 = ("The second input is " + strconv.Itoa(val2))
	return
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	_, c := vals()
	fmt.Println(c)

	info1, info2 := namedReturnValues(1, 2)
	fmt.Println(info1)
	fmt.Println(info2)

}
