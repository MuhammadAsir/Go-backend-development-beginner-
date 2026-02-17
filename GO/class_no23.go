package main

import "fmt"

func po(a int,b int,op func(x,y int)){ // higher order function.it will store function
  op(a,b)
}
func call() func(a int,b int){  // it returns a function
  return add
}

func add(a,b int){   //first order function. 
    fmt.Println(a+b)
}

func main(){
  add(10,12)
  po(5,7,add)  //po(4, 5, add) → এখানে add হলো callback function
  
  sum:=call()
  sum(1,2)
}
