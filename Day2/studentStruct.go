package main

import "fmt"

type StudentStruct struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	studentInformation := StudentStruct{
		firstName: "Don",
		lastName:  "Bosco",
		age:       23,
	}

	fmt.Print(" Student Name is : ", studentInformation.firstName+" "+studentInformation.lastName, " and Student age is : ", studentInformation.age)
}
