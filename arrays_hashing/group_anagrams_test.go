package arrays_hashing

import (
	"reflect"
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"groups",
			args{strs: []string{"act", "pots", "tops", "cat", "stop", "hat"}},
			[][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			"basic",
			args{strs: []string{"act", "pots", "tops", "cat", "stop", "hat"}},
			[][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := GroupAnagrams(tt.args.strs)
			if got := GroupAnagrams(tt.args.strs); !reflect.DeepEqual(sortSliceOfSlice(got), sortSliceOfSlice(tt.want)) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// this isn't very robust but whatever
func sortSliceOfSlice(ss [][]string) [][]string {
	for _, s := range ss {
		slices.Sort(s)
	}
	slices.SortFunc(ss, func(a, b []string) int {
		if len(a) < len(b) {
			return -1
		} else if len(a) > len(b) {
			return 1
		} else {
			if a[0] < b[0] {
				return -1
			} else {
				return 1
			}
		}
	})
	return ss
}
