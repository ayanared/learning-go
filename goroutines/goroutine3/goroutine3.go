package main

import (
	"fmt"
	"sync"
	"time"
)

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	//create a wait group - which basically holds the number of go routines you want your
	//program to wait for
	var wg sync.WaitGroup
	//The default number is 0, so we need to add a wait group
	wg.Add(1)

	//create an anonymous function in a goroutine
	go func() {
		count("sheep")
		//decrement the waitgroup counter
		wg.Done()
	}() //I don't know why these parenthesis are here...

	//tell the main function to wait till all the wait groups are finished

	wg.Wait()
}
