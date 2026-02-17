package main

import "fmt"

func prime(n int) bool {
	if n<=1{
		return false
	}
	for i := 2; i<n; i*=2 {
		if n%i == 0 {
			return false
		}
	}
	return true
	
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if prime(i) {
			fmt.Println(i, "is a prime number")
		}else{
			fmt.Println(i, "is not a prime number")
		}
	}
}
