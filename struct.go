package main

import "fmt"

type person struct {
	name string
	age int
}
func main(){
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice"})

	s := person{name: "sean",age:18}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age=11 // 指针的解引用
	fmt.Println(s.age)
}
