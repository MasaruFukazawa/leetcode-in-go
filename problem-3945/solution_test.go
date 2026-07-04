package main

import "testing"

func Test_digitFrequencyScore(t *testing.T) {

	datas := []struct {
		input  int
		output int
	}{
		{122, 5},
		{101, 2},
	}

	for _, data := range datas {
		if data.output != digitFrequencyScore(data.input) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
