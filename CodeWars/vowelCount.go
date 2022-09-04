package main

func GetCount(str string) (count int) {
	// Enter solution here\
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, c := range str {
		for _, v := range vowels {
			if v == c {
				count++
			}
		}
	}
	return count
}
