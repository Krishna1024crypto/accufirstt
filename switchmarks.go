package main
  
import "fmt"
  
func main() {
    var value int
    fmt.Scan(&value)
      
    // Switch statement without an     
    // optional statement and 
    // expression
   switch {
       case value >=90:
       fmt.Println("GRADE O")
       fmt.Println("OUTSTANDING")
       case value >=80:
       fmt.Println("GRADE A+")
       fmt.Println("EXCELLENT")
       case value >=70:
       fmt.Println("GRADE A")
       fmt.Println("GOOD")
       case value >=60:
       fmt.Println("GRADE B+")
       fmt.Println("BETTER")
       case value >=50:
       fmt.Println("GRADE B")
       fmt.Println("You passed")
       default: 
       fmt.Println("Invalid")
   }
  
}
