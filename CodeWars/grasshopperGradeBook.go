package main

func GetGrade(a, b, c int) rune {
	mean := (a + b + c) / 3
	var score rune
	if mean >= 90 && mean <= 100 {
		score = 'A'
	} else if mean >= 80 && mean < 90 {
		score = 'B'
	} else if mean >= 70 && mean < 80 {
		score = 'C'
	} else if mean >= 60 && mean < 70 {
		score = 'D'
	} else {
		score = 'F'
	}
	return score
}
