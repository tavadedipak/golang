package main

import "fmt"

type StudentStruct struct {
	firstName string
	lastName  string
	age       int8
}

func main() {

	studentInformation := StudentStruct{
		firstName: "Don",
		lastName:  "Bosco",
		age:       34,
	}

	fmt.Println(" Before Updaating Student Name is : ", studentInformation.firstName+" "+studentInformation.lastName, " and Student age is : ", studentInformation.age)

	studentInformation.updateStudentInformation()

	fmt.Print(" After Updaating Student Name is : ", studentInformation.firstName+" "+studentInformation.lastName, " and Student age is : ", studentInformation.age)
}

func (studentInformation *StudentStruct) updateStudentInformation() {

	*studentInformation = StudentStruct{
		firstName: "John",
		lastName:  "Bosco",
		age:       39,
	}

	fmt.Println("Updating Student Information")
}
