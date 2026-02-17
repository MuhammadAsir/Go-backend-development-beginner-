package main

import "fmt"
import "example.com/matlib"

// go mod init [name].com
func main(){
   fmt.Println("custom package: ")
	matlib.Add(4,7)
	
}