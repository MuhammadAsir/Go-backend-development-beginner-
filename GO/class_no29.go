package main

import "fmt"

func modifyArray(arr *[3]int) {
    arr[0] = 42
}


func main() {
	
	var arr=[3]int {1,2,3}
    fmt.Println(arr)
	for  i:=0;i<len(arr);i++ { //len means length of an array
		fmt.Println(arr[i])
	}
	modifyArray(&arr)
	cp:=arr  //copy of an array
	fmt.Println("Copy:", cp)

}
