package main

import "testing"

func Test_sumDivisibleByK(t *testing.T) {

	datas := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{1, 2, 2, 3, 3, 3, 3, 4}, 2, 16},
		{[]int{1, 2, 3, 4, 5}, 2, 0},
		{[]int{4, 4, 4, 1, 2, 3}, 3, 12},
	}

	for _, data := range datas {
		if data.output != sumDivisibleByK(data.nums, data.k) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
