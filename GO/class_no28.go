package main

import "fmt"

type USER struct {
	name string
	age  int
}

func (usr USER) Print() {
	fmt.Println(usr.name)
	fmt.Println(usr.age)
}

func main() {
	var user USER
	user = USER{
		name: "ASIR",
		age:  25,
	}
	user.Print()
}
