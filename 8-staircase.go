package main

import "fmt"

func staircase(n int32) {
	// Write your code here
	var output string
	var i int32
	var j int32
	var k int32
	for i = 1; i <= n; i++ {
		for j = n - 1; j >= i; j-- {
			output += " "
		}
		for k = 1; k <= i; k++ {
			output += "#"
		}
		output += "\n"
	}
	fmt.Println(output)
}
