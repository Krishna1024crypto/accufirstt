package main

import "fmt"
var sum int
var i int

func main() {
   /* local variable definition */
   var a int
   fmt.Scan(&a)
   if(a >1) {
   for i :=0; i<100 ; i++ {
   sum= sum+i+1; 
   fmt.Printf("sum: %d\n", sum)
   } } else if( a == 20 ) {
     
      fmt.Printf("Value of a is 20\n" )
   } else if( a == 30 ) {
      
      fmt.Printf("Value of a is 30\n" )
   } else {
      fmt.Printf("None of the values is matching\n" )
   }
   fmt.Printf("Exact value of a is: %d\n", a )
  }
