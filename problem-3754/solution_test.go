package main

import "testing"

func Test_sumAndMultiply(t *testing.T) {

	datas := []struct {
		input  int
		output int64
	}{
		{10203004, 12340},
		{1000, 1},
	}

	for _, data := range datas {
		if data.output != sumAndMultiply(data.input) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
