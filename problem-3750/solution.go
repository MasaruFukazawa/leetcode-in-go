package main

import "strconv"

func minimumFlips(n int) int {

	binary_n := strconv.FormatInt(int64(n), 2)
	counter := 0

	for i, j := 0, len(binary_n)-1; i < j; i, j = i+1, j-1 {
		if (binary_n[i] - '0') != (binary_n[j] - '0') {
			counter+=2
		}
    }

    return counter
}
