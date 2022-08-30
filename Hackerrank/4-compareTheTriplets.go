package main

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var alice int32 = 0
	var bob int32 = 0

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice += 1
		} else if a[i] < b[i] {
			bob += 1
		}
	}
	res := []int32{alice, bob}
	return res
}

// func main() {
// 	fmt.Printf("%v\n", compareTriplets([]int32{8, 2, 3}, []int32{3, 9, 10}))
// }
