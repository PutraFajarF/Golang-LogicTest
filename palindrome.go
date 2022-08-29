package main

import "fmt"

func main() {
	var words string
	fmt.Println("Input your words: ")
	fmt.Scan(&words)
	result := palindrome(words)
	if result == true {
		fmt.Println("This is Palindrome")
	} else {
		fmt.Println("This is not Palindrome")
	}
}

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
