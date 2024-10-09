package main

import "testing"

func Test_checkPalindrome(t *testing.T) {
	type args struct {
		word string
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
			if got := checkPalindrome(tt.args.word); got != tt.want {
				t.Errorf("checkPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
