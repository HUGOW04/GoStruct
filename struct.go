package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	me := person{
		name: "Hugo", age: 18,
	}
	fmt.Println(me.name, me.age)
}
