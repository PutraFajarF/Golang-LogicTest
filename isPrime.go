package main

import "fmt"

func isPrime(n int) int {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	fmt.Println(isPrime(7))
}
