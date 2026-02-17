package main

import "fmt"

func tower(n int, a int, b int, c int) {
	if n > 0 {
		tower(n-1, a, c, b)
		fmt.Println(a,c)
		tower(n-1, b, a, c)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	z := 1 << n
	fmt.Println(z-1)
	tower(n, 1, 2, 3)
}
