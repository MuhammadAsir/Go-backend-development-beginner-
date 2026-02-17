package main
import "fmt"
func main(){
	var a int
	fmt.Scan(&a)
	var arr1 =make([]int,a) 
     
	for i:=0;i<a;i++{
		fmt.Scan(&arr1[i])
	}

   fmt.Println(arr1)
}
