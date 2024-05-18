package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	result, err := Divide(10, 0) // change the second parameter

	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
