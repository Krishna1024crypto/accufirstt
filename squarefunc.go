package main

import "fmt"
func main() {
   var a int
   fmt.Scan(&a)
   var square int
   square=squ(a)
   fmt.Println(square)
   }
   func squ(n int ) int{
   return n*n
   }
   
   
