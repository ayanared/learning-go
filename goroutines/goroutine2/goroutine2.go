package main

import (
	"fmt"
	"time"
)

func count(thing string) {
	for i := 1; i < 100; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go count("deer")
	count("sheep")
	count("fish") //this will never run because sheep will never stop running... The solution is to put
	//in a goroutine (before the sheep)

}
