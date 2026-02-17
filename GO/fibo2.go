package main

import "fmt"

func printFibo(n int, a int, b int) {
	if n == 0 {
		return
	}
	fmt.Print(a, " ")
	printFibo(n-1, b, a+b)
}

func main() {
	var count int
	fmt.Print("Enter number of Fibonacci terms: ")
	fmt.Scan(&count)
	fmt.Println("Fibonacci series:")
	printFibo(count, 0, 1)
}
