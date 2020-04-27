package main

import "fmt"

func main() {
	studentRecord := make(map[int]string)
	studentRecord[01] = "Sachin Tendulkar"
	studentRecord[02] = "Virat Kohali"

	fmt.Println(" Before modifying Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

	fmt.Println("\n After deleting Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

	defer fmt.Println("\n Not able to delete Entries because defrered is used. Now Final Entry is : ", studentRecord)

	defer delete(studentRecord, 01)

}
