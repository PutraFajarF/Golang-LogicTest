package main

func RepeatStr(repetitions int, value string) string {
	var res string
	for i := 0; i < repetitions; i++ {
		res += value
	}
	return res
}
