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
	//then close the channel b/c the main channel is still waiting to recieve
	//the reciever should not ever try to close a channel
	close(c)
}

func main() {
	c := make(chan string)
	go count("sheep", c)
	//put in a for loop so that it will receive all the messages
	for {
		msg, open := <-c // recieve a second value - see if the channel still open

		if !open {
			break
		}
		fmt.Println(msg)
	}
	//could also use this syntactic sugar:
	//for msg := range c {
	//fmt.Println(msg)
	//}

}
