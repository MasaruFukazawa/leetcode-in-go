package main

import (
	"math"
	"strconv"
)

func reverse_num(s int) int {

	_s := strconv.Itoa(s)

	rs := []rune(_s)

	for i, j := 0, len(_s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	_return_rs, _ := strconv.Atoi(string(rs))

	return _return_rs
}

func mirrorDistance(n int) int {
	return int(math.Abs(float64(n - reverse_num(n))))
}
