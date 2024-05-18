package main

import "fmt"

func main() {
    var a int =10
    var b int =20
    // Using logical operators to evaluate conditions
    if a < b && b > 15 {
        fmt.Println("b is greater than a and greater than 15")
    }
    if a < 10 || b > 15 {
        fmt.Println("a is less than 10 or b is greater than 15")
    }
}
