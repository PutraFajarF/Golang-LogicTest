package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int32
	var counter int32
	var i int32
	for i = 0; i < int32(len(candles)); i++ {
		if candles[i] > max {
			max = candles[i]
			counter = 1
		} else if candles[i] == max {
			counter++
		}
	}
	return counter
}
