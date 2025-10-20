package main

func buildArray(nums []int) []int {
	return_nums := []int{}

	for _, v := range nums {
		return_nums = append(return_nums, nums[v])
	}

	return return_nums
}
