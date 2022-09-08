package main

import "strconv"

func StringToNumber(str string) int {
	a, _ := strconv.Atoi(str)
	return a
}
