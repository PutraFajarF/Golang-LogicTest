package main

import "fmt"

func fibonacci(n int) []int {
	n1 := 0
	n2 := 1
	var nextTerm int
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, n1)
		nextTerm = n1 + n2
		n1 = n2
		n2 = nextTerm
	}
	return result
}

func main() {
	fmt.Println(fibonacci(10))
}
