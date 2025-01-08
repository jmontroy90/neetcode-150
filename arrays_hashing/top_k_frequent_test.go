package arrays_hashing

import (
	"reflect"
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ex1", args{[]int{1, 2, 2, 3, 3, 3}, 2}, []int{2, 3}},
		{"ex2", args{[]int{7, 7}, 1}, []int{7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.args.nums, tt.args.k)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
