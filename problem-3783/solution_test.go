package main

import "testing"

func Test_mirrorDistance(t *testing.T) {

	datas := []struct {
		input  int
		output int
	}{
		{25, 27},
		{10, 9},
		{7, 0},
	}

	for _, data := range datas {
		if data.output != mirrorDistance(data.input) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
