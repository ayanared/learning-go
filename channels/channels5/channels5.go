package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) //create first channel of strings
	c2 := make(chan string) //create second channel of strings

	go func() { //create goroutine that sends the text every 1/2 second
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() { //create goroutine that sends the text every 2 seconds
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
