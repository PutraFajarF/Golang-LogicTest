package main

// Solution 1
func minimumSwaps(arr []int32) int32 {
	var count int32
	var i int32
	for i = 1; i < int32(len(arr)+1); i++ {
		if arr[i-1] != i {
			idx := findIndex(arr, i)
			arr[i-1], arr[idx] = arr[idx], arr[i-1]
			count++
		}
	}
	return count
}

func findIndex(arr []int32, val int32) int32 {
	for i, val1 := range arr {
		if val1 == val {
			return int32(i)
		}
	}
	return -1
}

// Solution 2
func minimumSwaps2(arr []int32) int32 {
	mapArr := make(map[int32]int32)
	for i, val := range arr {
		mapArr[val] = int32(i + 1)
	}

	var count int32
	var i int32
	for i = 1; i < int32(len(arr)+1); i++ {
		// 4:1 3:2 1:3 2:4 -- 4,3,1,2
		// 1:1 3:2 4:3 2:4 -- 1,3,4,2
		if mapArr[i] != i {
			arr[i-1], arr[mapArr[i]-1], mapArr[i], mapArr[arr[i-1]] = arr[mapArr[i]-1], arr[i-1], mapArr[arr[i-1]], mapArr[i]
			count++
		}
	}
	return count
}

// Solution 3
func minimumSwaps3(arr []int32) int32 {
	mapArr := make(map[int32]int32)
	for i, val := range arr {
		mapArr[val] = int32(i + 1)
	}

	var count int32
	var i int32
	for i = 1; i < int32(len(arr)+1); i++ {
		// 4:1 3:2 1:3 2:4 -- 4,3,1,2
		// 1:1 3:2 4:3 2:4 -- 1,3,4,2
		if mapArr[i] != i {
			mapArr[arr[i-1]] = mapArr[i]
			arr[mapArr[i]-1] = arr[i-1]
			count++
		}
	}
	return count
}
