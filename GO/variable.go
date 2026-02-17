package main
import "fmt"
var c = 3 // var Can be used inside and outside of functions
func main(){
  var student1 string="asir" // we can declare variable by var vrble name dtype=value
  fmt.Println(student1)
  a:="hossain"  // Use the := sign, followed by the variable value:
  fmt.Println(a)

  var is bool=true
  fmt.Println(is)

  var flt float32=10.99
  fmt.Println(flt)
  fmt.Println(c)
  const p=100 //it means we cant change the value of p
  fmt.Println(p)

}