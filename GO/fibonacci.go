package main
import "fmt"

func fibo(a int)int{
	if a<=2{
		return a-1
	}
	return fibo(a-1)+fibo(a-2)
}

func main(){
 
	var a int
	fmt.Scan(&a)
	n:=fibo(a)
	fmt.Println(n)
	
}