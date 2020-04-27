package main

import "fmt"

func main() {
	studentRecord := make(map[int]string)
	studentRecord[01] = "Virat Kohali"
	studentRecord[02] = "Sachin Tendulkar"

	fmt.Print(" Before modifying Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

	delete(studentRecord, 01)

	fmt.Print("\n After deleting Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

}
