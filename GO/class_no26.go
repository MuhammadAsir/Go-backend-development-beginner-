package main

import "fmt"
const n=10
var a=100

func outer(money int) func(){

age:=25
fmt.Println("age==",age)

show:=func(){
	money=money+n+a
	fmt.Println(money)
}
return show
}
func call(){
	inc1:=outer(100)
	inc1() //210
	inc1() //320
	inc1() //430
	incr2 := outer(100)
    incr2() //210
    incr2() //320
}

func main(){
	
 call()
}