package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int,error){
	if arg == 42 {
		return -1,errors.New("can't work with 42")
	}
	return arg+3,nil
}

func main(){
	res,err := f1(10)
	fmt.Println(res,err)
	res,err = f1(42)
	fmt.Println(res,err)
}
