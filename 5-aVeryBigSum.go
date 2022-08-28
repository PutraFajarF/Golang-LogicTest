package main

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var res int64 = 0
	for i := 0; i < len(ar); i++ {
		res += ar[i]
	}
	return res
}

// func main() {
// 	fmt.Printf("%v\n", aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
// }
