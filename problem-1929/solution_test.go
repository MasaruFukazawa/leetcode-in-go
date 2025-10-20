package main

import (
	"testing"
	"reflect"
)

func Test_getConcatenation(t *testing.T) {

	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}

	for _, tt := range tests {
		if !reflect.DeepEqual(getConcatenation(tt.input), tt.output) {
			t.Errorf("正解と計算結果が一致しません")
		}
	}
}
