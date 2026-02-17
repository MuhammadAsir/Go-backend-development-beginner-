package main
import "fmt"
 func main(){
  var age int     //declare variable
  fmt.Scan(&age)
  
  if age >= 18 {
	fmt.Println("Eligible")
  }else {             //  else is on the same line as }
	fmt.Println("NO")
  }

 }
    