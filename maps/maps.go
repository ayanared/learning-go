package main

import "fmt"

//Maps are Go's built-in associated data type (sometimes called hashes or dictionaries)
func main() {
	n := map[string]int{"foo": 1, "bar": 2} //initialize a map with values
	fmt.Println(n)
	m := make(map[string]int) //initialize a map
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)
	v1 := m["k1"]
	fmt.Println(v1)
	delete(m, "k2")
	fmt.Println(m)
}
