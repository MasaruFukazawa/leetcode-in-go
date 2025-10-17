package main

func sumDivisibleByK(nums []int, k int) int {

	return_value := 0
	nums_counter := make(map[int]int)

	for _, value := range nums {
		nums_counter[value] += 1
	}

	for key, value := range nums_counter {
		if 0 == (value % k) {
			return_value += (key * value)
		}
	}

	return return_value
}
