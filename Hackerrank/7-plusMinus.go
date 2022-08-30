package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var posCount int32 = 0
	var negCount int32 = 0
	var zeroCount int32 = 0
	var length int32 = int32(len(arr))
	for x := range arr {
		if arr[x] > 0 {
			posCount += 1
		} else if arr[x] < 0 {
			negCount += 1
		} else {
			zeroCount += 1
		}
	}
	fmt.Println(float64(posCount) / float64(length))
	fmt.Println(float64(negCount) / float64(length))
	fmt.Println(float64(zeroCount) / float64(length))
}
