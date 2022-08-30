package main

import "math"

func gradingStudents(grades []int32) []int32 {
	// Write your code here32
	var finalGrades []int32
	for i := range grades {
		quotient := float64(grades[i]) / float64(5)
		rounded := int32(math.Ceil(quotient) * 5)
		difference := rounded - grades[i]

		if grades[i] > 34 {
			if difference < 3 {
				finalGrades = append(finalGrades, rounded)
			} else {
				finalGrades = append(finalGrades, grades[i])
			}
		}

		if grades[i] < 35 {
			finalGrades = append(finalGrades, grades[i])
		}
	}
	return finalGrades
}
