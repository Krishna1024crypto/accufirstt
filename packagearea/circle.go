package main
   import(
   "fmt"
   "example.com/areacircle"
   )
   func main() {
   var area1, area2, area3 = areacircle.Area(2,4)
   fmt.Printf("Area of the circle is: %d", area1 )
   fmt.Printf("\nArea of the square is: %d", area2)
    fmt.Printf("\nArea of the rectangle is: %d", area3)
    }
