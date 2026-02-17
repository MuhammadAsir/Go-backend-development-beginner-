package main

import "fmt"

var (
    a = 20    //global variable
    b = 10
)

func add(x int, y int) {
    z := x + y
    fmt.Println(z)
}

func main() {
    p := 30     // local varible
    q := 40

    add(p, q)

    add(a, b)

    add(a, p)

    
}

