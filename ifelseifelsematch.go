package main

import "fmt"

func main() {
   var a int
   fmt.Scan(&a)
   if( a == 100 ) {
      
      fmt.Printf("Value of a is 10\n" )
   } else if( a == 20 ) {
      /
      fmt.Printf("Value of a is 20\n" )
   } else if( a == 30 ) {
    
      fmt.Printf("Value of a is 30\n" )
   } else {
      
      fmt.Printf("None of the values is matching\n" )
   }
   fmt.Printf("Exact value of a is: %d\n", a )
  }
