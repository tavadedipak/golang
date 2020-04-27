package main

import "fmt"

func main() {
	myTechnologySlice := []string{"Docker", "Linux", "Python", "Cloud", "Windows", "Mac"}
	
	for _, technologies := range myTechnologySlice {
		fmt.Print("\n My Techology are : ", technologies)
	}
	
	fmt.Print("\n Lenth of the slice is ", len(myTechnologySlice))
}
