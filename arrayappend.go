 package main
  
import (
    "fmt"
)
  
func main() {
  
    
    arraynum := []int{234, 567, 7890, 1234, 234}
    arraystring := []string{"APPLE", "MANGO", "BANANA"}
  
    
    fmt.Println("Array_1: ", arraynum)
    fmt.Println("Length of Array_1:", len(arraynum))
    fmt.Println("Capacity of Array_1: ", cap(arraynum))
    fmt.Println("Array_2: ",arraystring)
    fmt.Println("Length of Array_2: ", len(arraystring))
    fmt.Println("Capacity of Slice_2: ", cap(arraystring))
    new1 := append(arraynum, 1000)
    new2 := append(arraystring, "GUAVA")
    fmt.Println("New : ", new1)
    fmt.Println("New length : ", len(new1))
    fmt.Println("New capacity:", cap(new1))
    fmt.Println("New:", new2)
    fmt.Println("New length of new2 ", len(new2))
    fmt.Println("New capacity of new_2: ", cap(new2))
}
