package two_pointers

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "ex1", args: args{numbers: []int{-1, 0, 1, 2, -1, -4}}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "ex2", args: args{numbers: []int{-1, -1, 2, 0, 0, -2, 3, 1}}, want: [][]int{{-2, -1, 3}, {-1, -1, 2}, {-1, 0, 1}}},
		{name: "ex3", args: args{numbers: []int{0, 1, 1}}, want: nil},
		{name: "ex4", args: args{numbers: []int{0, 0, 0}}, want: [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
