package arrays_hashing

import "testing"

func TestIsValidSudoku(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{[][]int{
			{1, 2, 0, 0, 3, 0, 0, 0, 0},
			{4, 0, 0, 5, 0, 0, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 0, 3},
			{5, 0, 0, 0, 6, 0, 0, 0, 4},
			{0, 0, 0, 8, 0, 3, 0, 0, 5},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 0, 0, 0, 0, 0, 2, 0, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 8},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		}}, true},
		{"invalid", args{[][]int{
			{1, 2, 0, 0, 3, 0, 0, 0, 0},
			{4, 0, 0, 5, 0, 0, 0, 0, 0},
			{0, 9, 1, 0, 0, 0, 0, 0, 3},
			{5, 0, 0, 0, 6, 0, 0, 0, 4},
			{0, 0, 0, 8, 0, 3, 0, 0, 5},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 0, 0, 0, 0, 0, 2, 0, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 8},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("IsValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
