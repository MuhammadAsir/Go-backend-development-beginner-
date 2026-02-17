package main

import "fmt"

func main() {
    // Create an array of strings
    arr := [6]string{"This", "is", "a", "Go", "interview", "Questions"}
    fmt.Println(arr)

    // Create a slice from array indexes 1 to 3 (exclusive of 4)
    s := arr[1:4]
    fmt.Println(s) // [is a Go]

    // Create a slice from a slice
    s1 := s[1:2]
    fmt.Println(s1) // [a]
    fmt.Println(len(s1)) // 1
    fmt.Println(cap(s1)) // 4 (capacity depends on the underlying array)

    // Slice literal
    s2 := []int{3, 4, 7}
    fmt.Println("slice", s2, "lenght:", len(s2), "capacity:", cap(s2))

    // make() function with length only
    s3 := make([]int, 3)
    s3[0] = 5
    fmt.Println(s3)
    fmt.Println(len(s3))
    fmt.Println(cap(s3))

    // make() function with length and capacity
    s4 := make([]int, 3, 5)
    s4[0] = 5
    fmt.Println(s4)
    fmt.Println(len(s4))
    fmt.Println(cap(s4))

    // Empty slice
    var s5 []int
    fmt.Println(s5) // []

    // Appending elements to empty slice
    var s6 []int
    s6 = append(s6, 1)
    fmt.Println(s6) // [1]

    var s7 []int
    s7 = append(s7, 1, 2, 3)
    fmt.Println(s7, len(s7), cap(s7)) // [1 2 3] 3 3

    // Interview question: Sharing underlying array
    var x []int
    x = append(x, 1)
    x = append(x, 2)
    x = append(x, 3)

    y := x
    x = append(x, 4)
    y = append(y, 5)

    x[0] = 10

    fmt.Println(x) // [10 2 3 5]
    fmt.Println(y) // [10 2 3 5]

    // Another interview question
    slc := []int{1, 2, 3, 4, 5}
    slc = append(slc, 6)
    slc = append(slc, 7)

    slcA := slc[4:]

    slcY := changeSlice(slcA)

    fmt.Println(slc)  // [1 2 3 4 10 6 7]
    fmt.Println(slcY) // [10 6 7 11]
    fmt.Println(slc[0:8]) // [1 2 3 4 10 6 7 11]

    // Variadic function call
    variadic(2, 3, 4, 6, 8, 10)
}

// Function that changes the slice passed
func changeSlice(a []int) []int {
    a[0] = 10
    a = append(a, 11)
    return a
}

// Variadic function that takes multiple integers
func variadic(numbers ...int) {
    fmt.Println(numbers)
    fmt.Println(len(numbers))
    fmt.Println(cap(numbers))
}
