package main

import (
	"strconv"
)

func digitFrequencyScore(n int) int {

	sum := 0

	s := strconv.Itoa(n)

	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
	}

	return sum
}
