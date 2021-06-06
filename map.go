package main

import "fmt"

func main() {   
  
   relation :=map[string]string{"COMPANY":"EMPLOYEES","COLLEGE":"STUDENTS","COUNTRY":"CITIZENS"}
   relation["HUMANITY"]= "PEOPLES"
   fmt.Println(relation)   
   value_name, ok := relation["COMPANY"]
    fmt.Println("\nKey present or not:", ok)
    fmt.Println("Value:", value_name) 
    relation2 := relation
    delete(relation2,"COMPANY")
    fmt.Println(relation2)} 
    
