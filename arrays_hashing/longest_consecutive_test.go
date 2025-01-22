package arrays_hashing

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{[]int{2, 20, 4, 10, 3, 4, 5}}, 4},
		{"7", args{[]int{0, 3, 2, 5, 4, 6, 1, 1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("LongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestConsecutiveHash(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{[]int{2, 20, 4, 10, 3, 4, 5}}, 4},
		{"7", args{[]int{0, 3, 2, 5, 4, 6, 1, 1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutiveHash(tt.args.nums); got != tt.want {
				t.Errorf("LongestConsecutiveHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
