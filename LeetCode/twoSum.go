package main

func twoSum(nums []int, target int) []int {
	// var res []int
	// for i:= 0; i < len(nums); i++ {
	//     for j:= 1 + i; i < len(nums); i++{
	//         if nums[i]+nums[j] == target {
	//             res = append(res, i, j)
	//         }
	//     }
	// }
	// return res
	record := make(map[int]int)

	for index, num := range nums {
		difference := target - num
		if res, ok := record[difference]; ok {
			return []int{index, res}
		}
		record[num] = index
	}

	return []int{}
}
