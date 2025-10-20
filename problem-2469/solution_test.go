package main

import (
	"reflect"
	"testing"
)

func Test_convertTemperature(t *testing.T) {

	tests := []struct {
		input  float64
		output []float64
	}{
		{36.50, []float64{309.65000, 97.70000}},
		{122.11, []float64{395.26000, 251.79800}},
	}

	for _, tt := range tests {
		got := convertTemperature(tt.input)
		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
