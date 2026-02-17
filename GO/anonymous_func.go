// anonymous function
// IIFE - ইমিডিয়েটলি ইনভোকড function expression

package main

import "fmt"

func main() {
    // anonymous function
    func(a int, b int) {
        c := a + b
        fmt.Println(c)
    }(5, 7) // IIFE  
}

func init() {
    fmt.Println("I am")
}
