package main

import "fmt"

func ping(ping chan<- string, msg string){
	ping <- msg
}
func pang(ping <-chan string, pang chan<- string){
	msg := <- ping
	pang <- msg
}
func main(){
	pings := make(chan string,1)
	pangs := make(chan string,1)
	ping(pings,"msg")
	pang(pings,pangs)
	fmt.Println(<-pangs)
}
