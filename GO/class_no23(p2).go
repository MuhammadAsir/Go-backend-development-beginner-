package main

import "fmt"

// higher-order function: takes a function and can also returns a function
func po(a int, b int, op func(x, y int))func(a int, b int) {
	op(a, b) // call the function once
	return add // returning the function
}

// first-order function
func add(a, b int) {
	fmt.Println(a + b)
}

func call() func(a int,b int){  // it returns a function
  return add
}
func main() {
	add(10, 12)        // prints 22
	po(5, 7, add)      // prints 12

	anonAdd := func() {  //anonymous function
		fmt.Println("NOOO")
	}
	anonAdd()

	sum := po(1,2,add) // prints 3
    sum(5,6)// prints 11
  
  AAA:=call()
  AAA(4,3)

}

func init() {
	fmt.Println("YESSS")
}
