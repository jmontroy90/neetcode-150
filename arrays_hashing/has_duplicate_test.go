package arrays_hashing

import (
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no-dupe-ordered", args{nums: []int{1, 2, 3, 4, 5}}, false},
		{"no-dupe-unordered", args{nums: []int{1, 3, 4, 6, 5}}, false},
		{"dupe-ordered", args{nums: []int{1, 2, 3, 3, 4}}, true},
		{"dupe-unordered", args{nums: []int{1, 3, 2, 3, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("HasDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicateHash(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no-dupe-ordered", args{nums: []int{1, 2, 3, 4, 5}}, false},
		{"no-dupe-unordered", args{nums: []int{1, 3, 4, 6, 5}}, false},
		{"dupe-ordered", args{nums: []int{1, 2, 3, 3, 4}}, true},
		{"dupe-unordered", args{nums: []int{1, 3, 2, 3, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicateHash(tt.args.nums); got != tt.want {
				t.Errorf("HasDuplicateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
