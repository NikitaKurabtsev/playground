package main

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Already sorted array",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reverse sorted array",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Array with duplicate elements",
			args: args{arr: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}},
			want: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name: "Empty array",
			args: args{arr: []int{}},
			want: []int{},
		},
		{
			name: "Array with one element",
			args: args{arr: []int{1}},
			want: []int{1},
		},
		{
			name: "Array with negative numbers",
			args: args{arr: []int{-3, 0, 5, -2, 1, -4}},
			want: []int{-4, -3, -2, 0, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
