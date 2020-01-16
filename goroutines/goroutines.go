package main

import (
	"fmt"
	"time"
)

// a goroutine is a lightweight thread managed by the Go runtime

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//the function, say, is called in the "main" thread of execution
	say("hello")
	//putting go before the function, say, makes this function run in another thread (besides main)
	go say("world")

}
