package main

import (
	"fmt"
	"time"
)

var message1 = make(chan string, 10)

func input() {
	message1 <- "hello0"
	message1 <- "hello1"
	message1 <- "hello2"
	message1 <- "hello3"
	message1 <- "hello4"
	message1 <- "hello5"
	message1 <- "hello6"
	message1 <- "hello7"
	message1 <- "hello8"
	message1 <- "hello9"
}
func input1() {
	message1 <- "hello10"
	message1 <- "hello11"
	message1 <- "hello12"

}
func main() {
	go input()
	time.Sleep(time.Second * 4)
	var i = 0
	for str := range message1 {
		i++
		fmt.Println(str)
		if i == 4 {
			time.Sleep(time.Second * 1)
			go input1()
		}
	}

}
