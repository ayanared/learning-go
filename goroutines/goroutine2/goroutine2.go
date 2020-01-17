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
	count("fish") //this won't run till after the sheep stop running... The solution is to put
	//in a goroutine (before the sheep) - then it will run at the same time

}
