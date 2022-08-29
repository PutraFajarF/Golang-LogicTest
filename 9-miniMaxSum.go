package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	var min int64
	var max int64
	for i := 0; i < len(arr); i++ {
		if sum-int64(arr[i]) <= min || min == 0 {
			min = sum - int64(arr[i])
		}
		if sum-int64(arr[i]) > max || max == 0 {
			max = sum - int64(arr[i])
		}
	}
	fmt.Printf("%d %d", min, max)
}
