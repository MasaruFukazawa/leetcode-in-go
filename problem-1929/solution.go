package main

func getConcatenation(nums []int) []int {
	// ... で 配列を展開してくれる
	return append(nums, nums...)
}
