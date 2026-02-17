package main

import "fmt"

// গ্লোবাল ফাংশন এক্সপ্রেশন
var add = func(x, y int) {
    fmt.Println(x + y)
}

func main() {
    add(4, 7) // গ্লোবাল `add` কল হচ্ছে

    // লোকাল ভেরিয়েবলে ফাংশন এক্সপ্রেশন অ্যাসাইন করা
    add := func(a int, b int) {
        c := a + b
        fmt.Println(c) 
    }

    add(2, 3) // এখন লোকাল `add` কল হচ্ছে
}

func init() {
    fmt.Println("আমি প্রথমে কল হব")
}
