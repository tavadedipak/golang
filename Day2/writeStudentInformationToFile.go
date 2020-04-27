package main

import (
	"encoding/json"
	"io/ioutil"
)

type StudentStruct struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	studentInformation := StudentStruct{
		firstName: "Don",
		lastName:  "Bosco",
		age:       36,
	}

	writefile, _ := json.MarshalIndent(studentInformation, "", " ")

	_ = ioutil.WriteFile("test.json", writefile, 0644)

}
