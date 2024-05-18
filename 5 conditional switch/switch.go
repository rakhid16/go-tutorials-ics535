package main

import "fmt"

func main() {
    grade := "B"
    switch grade {
    case "A":
        fmt.Println("Excellent!")
    case "B", "C":
        fmt.Println("Well done")
    case "D":
        fmt.Println("You passed")
    case "F":
        fmt.Println("Better try again")
    default:
        fmt.Println("Invalid grade")
    }
}
