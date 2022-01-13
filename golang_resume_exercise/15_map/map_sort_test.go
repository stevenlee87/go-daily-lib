package main

import "testing"

func Test_isSubset(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubset(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
