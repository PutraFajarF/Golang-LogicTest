package main

import "math"

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	diagonal1 := []int32{}
	diagonal2 := []int32{}
	var sumDiag1 int32 = 0
	var sumDiag2 int32 = 0
	for x := range arr {
		for y := range arr {
			if x == y {
				diagonal1 = append(diagonal1, arr[x][y])
				sumDiag1 += arr[x][y]
			}
			if x == len(arr)-y-1 {
				diagonal2 = append(diagonal2, arr[x][y])
				sumDiag2 += arr[x][y]
			}
		}
	}
	return int32(math.Abs(float64(sumDiag1) - float64(sumDiag2)))
}

// Another Solution
func diagonalDifference2(arr [][]int32) int32 {
	// Write your code here
	var sumDiag1 int32 = 0
	var sumDiag2 int32 = 0
	for x := range arr {
		for y := range arr {
			if x == y {
				sumDiag1 += arr[x][y]
			}
			if x+y == len(arr)-1 {
				sumDiag2 += arr[x][y]
			}
		}
	}
	return int32(math.Abs(float64(sumDiag1 - sumDiag2)))
}
