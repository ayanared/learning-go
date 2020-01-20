package main

import (
	"fmt"
	"time"
)

// use the blocking nature of channels to syncronize go routines.
//
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// send the value of thing over the channel (arrow pointing into means to send a message)
		c <- thing //when you send a message, it will wait till the receiver is ready to receive
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	c := make(chan string)
	go count("sheep", c)
	//send the value into the var, msg
	msg := <-c //have to wait for there to be a value to receive
	fmt.Println(msg)

}
