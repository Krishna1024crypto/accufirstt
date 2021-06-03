package main  
  
import (  
 "fmt"  
)  
  
func main() {  
 var no1, no2, no3 int  
 fmt.Print("Enter First No:")  
 fmt.Scanln(&no1)  
 fmt.Print("Enter Second Number:")  
 fmt.Scanln(&no2)  
 fmt.Print("Enter Third Number:")  
 fmt.Scanln(&no3)  
  
 fmt.Printf("The Entered three numbers are %d %d %d \n", no1, no2, no3)  
  
 if number1 >= no2 && no1 >= no3 {  
  fmt.Println("Largest of three numbers: ", no1)  
 } else if number2 >= number1 && number2 >= no3 {  
  fmt.Println("Largest of three numbers: ", no2)  
 } else {  
  fmt.Println("Largest of three numbers: ", no3)  
 }  
}  