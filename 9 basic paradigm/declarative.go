package main

import "fmt"

type Address struct {
	Street_block string
	Zip_code     int
	City         string
	Country      string
}

type Student struct {
	Name    string
	Id      int
	Address Address
}

func main() {
	student_1 := Student{
		Name: "Giorno Giovana",
		Id:   567,
		Address: Address{
			Street_block: "Golden street number 5",
			Zip_code:     1234,
			City:         "Naples",
			Country:      "Italy",
		},
	}

	fmt.Println("Student details:", student_1)
}
