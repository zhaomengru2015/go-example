package main

import "fmt"

func zeroval(ival int){
	ival =0
}
func zeroptr(iptr *int){
	*iptr = 0
}
func main(){
	i :=7
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
