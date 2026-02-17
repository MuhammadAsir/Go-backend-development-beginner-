package main

import "fmt"

func print(a int, b int)(int ,int) { //(int,int) will return two values
	
	sum:=a+b
	mul:=a*b
	return sum,mul
}

func main() {
	a:=10
	b:=20
   
	fmt.Println(print(a,b))


}
