package main   

// func init must have no arguments and no return values and we cant call this func 
// init func wil execute first 
import "fmt"

var a=10

func main(){

	fmt.Println(a)
	

}
func init(){
	
	fmt.Println(a)
	a=20

}
