package main

import (
	"strconv"
)

func countSheep(num int) string {
	// Your code here!
	var a string
	for i := 1; i <= num; i++ {
		a += strconv.Itoa(i) + " sheep..."
	}
	return a
}
