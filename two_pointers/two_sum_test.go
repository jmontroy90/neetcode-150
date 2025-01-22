package two_pointers

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"3", args{[]int{1, 2, 3, 4}, 3}, []int{0, 1}},
		{"5", args{[]int{1, 2, 4, 5}, 5}, []int{0, 2}},
		{"9", args{[]int{2, 3, 4, 7, 11, 15}, 9}, []int{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumPointers(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"3", args{[]int{1, 2, 3, 4}, 3}, []int{0, 1}},
		{"5", args{[]int{1, 2, 4, 5}, 5}, []int{0, 2}},
		{"9", args{[]int{2, 3, 4, 7, 11, 15}, 9}, []int{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumPointers(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
