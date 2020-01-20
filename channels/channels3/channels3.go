package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) // make buffered channel of strings with capacity
	// without capacity, send will block until
	// something is ready to recieve (in another go routine)
	// the next line wont get executed
	// it won't block until the buffer is full
	c <- "hello" //send hello across channel

	msg := <-c // recieve the string from the channel and put it in the variable msg
	fmt.Println(msg)
}
