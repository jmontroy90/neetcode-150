package arrays_hashing

import (
	"reflect"
	"testing"
)

func TestProductExceptSelfDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ex1", args{[]int{1, 2, 4, 6}}, []int{48, 24, 12, 8}},
		{"ex2", args{[]int{-1, 0, 1, 2, 3}}, []int{0, -6, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelfDivision(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelfDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelfPrefixSuffix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ex1", args{[]int{1, 2, 4, 6}}, []int{48, 24, 12, 8}},
		{"ex2", args{[]int{-1, 0, 1, 2, 3}}, []int{0, -6, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelfPrefixSuffix(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelfPrefixSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
