package main

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {

	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}

	for _, tt := range tests {
		got := buildArray(tt.input)
		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
