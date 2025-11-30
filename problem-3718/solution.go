package main

func missingMultiple(nums []int, k int) int {

	return_value := 0
	quotients := make(map[int]struct{})

	for _, num := range nums {
		if (num % k) == 0 {
			quotients[num/k] = struct{}{}
		}
	}

	for i := 1; i <= 101; i++ {
		if _, exists := quotients[i]; !exists {
			return_value = i * k
			break
		}
	}

	return return_value
}
