package main

import "fmt"

func main() {
   
   numbers := []int{0,1,2,3,4,5,6,7,8,9,10}   
   fmt.Println(numbers)
   
  
   fmt.Println("numbers ==", numbers)
   
   /* Slice is used to print desired elements*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])
   fmt.Println("numbers[:4] ==", numbers[:4])
   fmt.Println("numbers[5:] ==", numbers[5:])
   }
