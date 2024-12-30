package main

import "testing"

func Test_missingNumberSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"test2", args{[]int{3, 0, 1}}, 2},
		{"test3", args{[]int{0, 1}}, 2},
		{"test4", args{[]int{3, 2, 4, 6, 0, 1}}, 5},
		{"test5", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberSum(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
