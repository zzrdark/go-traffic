package main

import (
	"fmt"
	"strconv"
	"time"
)


 var messageString = make (chan string,100)
 var messageInteger = make (chan int,100)


func inputString(){
	for i := 0 ; i<10 ; i++{
		messageString <- "hello String："+ strconv.Itoa(i)
		time.Sleep(time.Second)
	}
}
func inputInteger(){
	for i := 0 ; i<10 ; i++{
		messageInteger <- i
		time.Sleep(time.Second)
	}
}

func main() {

	//for i:= 0 ;i<10;i++{
		go inputInteger()
		go inputString()
	//}
	for {
		select {
		case p1,strCheck := <- messageString :
			if !strCheck{
				fmt.Println("strEnd")
			}
			fmt.Println(p1)
		case p2,intCheck := <- messageInteger :
			if !intCheck{
				fmt.Println("intEnd")
			}
			fmt.Println(p2)
		}
	}




	time.Sleep(time.Second*20)
}
