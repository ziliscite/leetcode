package main

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbersBrute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"test2", args{[]int{1, 1}}, []int{2}},
		{"test3", args{[]int{1, 2, 2}}, []int{3}},
		{"test4", args{[]int{1, 1, 1, 1, 2, 3}}, []int{4, 5, 6}},
		{"test5", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbersBrute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbersBrute() = %v, want %v", got, tt.want)
			}
		})
	}
}
