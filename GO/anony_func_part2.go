// anonymous function
// IIFE - ইমিডিয়েটলি ইনভোকড function expression

package main

import "fmt"

func main() {
    // anonymous function
    func() {
        c := 1 + 2
        fmt.Println(c)
    }() // IIFE  
}

func init() {
    fmt.Println("I am")
}
