package main

import (
	"strconv"
	"strings"
)

func sumAndMultiply(n int) int64 {

	sum := 0
	non_zero_digit_str := strings.ReplaceAll(strconv.Itoa(n), "0", "")
	non_zero_digit, _ := strconv.Atoi(non_zero_digit_str)

	for _, char := range non_zero_digit_str {
		sum += int(char - '0')
	}

	return int64(non_zero_digit * sum)
}
