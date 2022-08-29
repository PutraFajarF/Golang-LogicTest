package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	// input: 07:15:45AM or 11:03:33PM
	s = strings.ToLower(s)
	pm := strings.HasSuffix(s, "pm")
	am := strings.HasSuffix(s, "am")

	t := s[:len(s)-2]

	timeArr := strings.Split(t, ":")
	hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]
	hhInt, err := strconv.Atoi(hh) // string to int

	if err != nil {
		panic(err.Error())
	}

	if am && hhInt == 12 {
		hhInt = 0
	}

	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt) // int to string

	// output: 07:15:45 or 23:03:33
	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)
}
