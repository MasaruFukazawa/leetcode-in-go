package main

import (
	"testing"
)

func Test_defangIPaddr(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"pattern_01", "1.1.1.1", "1[.]1[.]1[.]1"},
		{"pattern_02", "255.100.50.0", "255[.]100[.]50[.]0"},
	}

	for _, tt := range tests {

		got := defangIPaddr(tt.input)

		t.Run(tt.name, func(t *testing.T) {
			if got != tt.output {
				t.Errorf("got %v, want %v", got, tt.output)
			}
		})
	}
}
