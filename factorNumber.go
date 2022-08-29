package main

import "fmt"

// func main() {
// 	var factorSum int
// 	fmt.Print("Type Number: ")
// 	fmt.Scanln(&factorSum)

// 	fmt.Println("Factorial Number: ", factorSum, "=")
// 	findFactorNum(factorSum)
// }

func findFactorNum(factorSum int) {
	for i := 1; i <= factorSum; i++ {
		if factorSum%i == 0 {
			fmt.Println(i)
		}
	}
}
