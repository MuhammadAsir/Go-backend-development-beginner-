package main

import "fmt"

func sum(){
   add(2,6) 
}
func add(a int,b int){
    c:=a+b
    fmt.Println(c)
}

func main() {
    sum()
    add(2,2)

    add:=func(a int,b int){
        c:=a+b
    fmt.Println(c)
    }
   add(5,5)
 
}

func init() {
    fmt.Println("I am")
}
