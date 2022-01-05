package main

import "fmt"

func main() {
	str := "This is string ops example !"

	// := is used only for the first time declaration and assignment
	//str = str + "!!"
	//fmt.Println(str)
	fmt.Printf("\nValue of str : %v and Type of str is : %T \n", str, str)

	// Substring or slicing of a String
	fmt.Println(str[4:10])
	fmt.Println(str[5:])

	// Multiline Stirng
	poem := `
			Johny Johny, Yes Papa
			Drinking Beer ? No Papa !
			Telling Lies ? Yes Papa !
	`

	fmt.Println(poem)
}
