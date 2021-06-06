package main
import "fmt"
type person struct{
   age int
   name string
   address string
   }
   func main(){
   person1 :=new(person)
   fmt.Println("Enter the age")
   fmt.Scan(&person1.age)
   fmt.Println("Enter the name")
   fmt.Scan(&person1.name)
   fmt.Println("Enter the adress")
   fmt.Scan(&person1.address)
   fmt.Println("The details are:",person1.age,person1.name,person1.address)}
