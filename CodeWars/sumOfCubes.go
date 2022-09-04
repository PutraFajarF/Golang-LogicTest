package main

func SumCubes(n int) int {
	// your code here
	var a int
	for i := 1; i <= n; i++ {
		a += i * i * i
	}
	return a
}
