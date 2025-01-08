package arrays_hashing

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one", args{[]int{3, 4, 5, 6}, 7}, []int{0, 1}},
		{"two", args{[]int{4, 5, 6}, 10}, []int{0, 2}},
		{"three", args{[]int{5, 5}, 10}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumHash(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one", args{[]int{3, 4, 5, 6}, 7}, []int{0, 1}},
		{"two", args{[]int{4, 5, 6}, 10}, []int{0, 2}},
		{"three", args{[]int{5, 5}, 10}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumHash(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
