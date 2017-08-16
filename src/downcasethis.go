package main

import "fmt"
import "strings"

func main() {
	var input string
	fmt.Println("---DOWNCASETHIS---")
	fmt.Printf("Input a string you would like to downcase -->  ")
	fmt.Scanf("%s", &input)
	fmt.Printf("Your downcased string is ")
	fmt.Printf(strings.ToLower(input))
}