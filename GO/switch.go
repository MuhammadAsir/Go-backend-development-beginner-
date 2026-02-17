package main
import "fmt"
func main(){
var a int
fmt.Scan(&a)

switch a{
case 1:
	fmt.Println("ASIR")
case 2,3:
	fmt.Println("NO")
default:
	fmt.Println("ERROR")
}

}