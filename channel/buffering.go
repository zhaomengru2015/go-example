package main

import "fmt"

func main(){
	message := make(chan string, 2)
	go func(){message <- "world"}()
	message <- "world2"
	msg1 := <- message
	msg2 := <- message
	fmt.Println(msg1, msg2)
}
