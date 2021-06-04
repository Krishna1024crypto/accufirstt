package main
import "fmt"

func main() {
var a=10
var b=20
var result int
result = max(a,b)
fmt.Println("max value is: %d",result)	
}
func max(n1, n2 int) int{
if(n1>n2){
var result int =n1
return result
}else{
var result int =n2
return result
}
}


