package main

import "testing"

func Test_theMaximumAchievableX(t *testing.T) {

	datas := []struct {
		num  int
		t  int
		output int
	}{
		{4, 1, 6},
		{3, 2, 7},
	}

	for _, data := range datas {
		if data.output != theMaximumAchievableX(data.num, data.t) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
