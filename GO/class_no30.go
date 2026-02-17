package main

import "fmt"

type User struct {
    Name   string
    Age    int
    Salary float64
}

func p1(numbers [3]int) {
    fmt.Println(numbers)
}

func swap(x, y *int) {
    temp := *x
    *x = *y
    *y = temp
}

func main() {
    x := 20
    p := &x
    *p = 30

    fmt.Println(x)           // 30
    fmt.Println("Address:", p)
    fmt.Println("Value:", *p)

    arr := [3]int{1, 2, 3}
    p1(arr)   // pass by value
	
	a, b := 5, 6
    swap(&a, &b)
    fmt.Println(a, b) // 2 1

    user1 := User{
        Name:   "Ruhin",
        Age:    21,
        Salary: 0,
    }
    p2 := &user1
    fmt.Println(p2.Age)
}
