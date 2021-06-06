package main
  
import "fmt"
func myfunc(p, q int)( circle int, square int,rectangle int ){
    rectangle = p*q
    square = p*p
    circle=22/7 *p*p
    return  
}
  
func main() {
      
    
   var area1, area2, area3 = myfunc(2,4)
     
   // Display the values
   fmt.Printf("Area of the circle is: %d", area1 )
   fmt.Printf("\nArea of the square is: %d", area2)
    fmt.Printf("\nArea of the rectangle is: %d", area3)
     
}
