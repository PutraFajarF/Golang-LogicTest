package main

import "fmt"

/*
 * Complete the 'reverseArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func reverseArray(a []int32) []int32 {
	// Write your code here
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	fmt.Printf("%v\n", reverseArray([]int32{8, 2, 3, 9, 10}))
	fmt.Printf("%v\n", reverseArray([]int32{1, 2, 3, 4}))
}
