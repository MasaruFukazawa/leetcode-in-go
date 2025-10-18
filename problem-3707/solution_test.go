package main

import "testing"

func Test_scoreBalance(t *testing.T) {

	datas := []struct {
		input  string
		output bool
	}{
		{"adcb", true},
		{"bace", false},
	}

	for _, data := range datas {
		if data.output != scoreBalance(data.input) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
