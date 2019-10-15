package main

import (
	"fmt"
	"time"
)

var message = make(chan string)

const str = 1

func sample() {
	message <- "hello"
}

func sample1() {
	time.Sleep(time.Second * 2)
	var str = <-message
	str = str + " too"
	message <- str
}

func main() {

	go sample()
	go sample1()
	time.Sleep(time.Second * 3)
	fmt.Println(<-message)
}
