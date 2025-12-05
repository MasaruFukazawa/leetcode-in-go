package main

import "testing"

func Test_minimumFlips(t *testing.T) {

	datas := []struct {
		name   string
		input  int
		output int
	}{
		{"ケース1", 7, 0},   // 111 - 回文
		{"ケース2", 10, 4},  // 1010 - 完全非回文
		{"ケース3", 5, 0},   // 101 - 奇数長の回文
		{"ケース4", 9, 0},   // 1001 - 偶数長の回文
		{"ケース5", 6, 2},   // 110 - 部分的に異なる (位置0と2が異なる)
	}

	for _, data := range datas {
		if data.output != minimumFlips(data.input) {
			t.Errorf("%s : 正解と計算結果が一致しません", data.name)
		}
	}
}
