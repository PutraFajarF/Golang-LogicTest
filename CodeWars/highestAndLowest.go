package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// best solution
func HighAndLow(in string) string {
	// Code here or
	numStrings := strings.Fields(in)
	var nums = []int{}

	for _, i := range numStrings {
		j, _ := strconv.Atoi(i)
		nums = append(nums, j)
	}
	sort.Ints(nums)
	return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}

// good solution
func HighAndLow2(in string) string {
	var tmpH, tmpL int
	for i, s := range strings.Fields(in) {
		n, _ := strconv.Atoi(string(s))
		if i == 0 {
			tmpH = n
			tmpL = n
		}

		if n > tmpH {
			tmpH = n
		}

		if n < tmpL {
			tmpL = n
		}
	}
	return fmt.Sprintf("%d %d", tmpH, tmpL)
}

// another solution
func HighAndLow3(in string) string {
	// divide the string into little bitty pieces
	tokens := strings.Split(in, " ")

	// convert the slice from strings to integers
	nums := make([]int, 0, len(tokens))
	for _, t := range tokens {
		// attempt to convert alpha to integer
		n, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}

		// add the newly converted integer to the slice
		nums = append(nums, n)
	}

	// this sorts them in-place, destroying the previous order
	sort.Ints(nums)

	// return our highest (end of the slice), and our lowest (beginning of the slice)
	return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}
