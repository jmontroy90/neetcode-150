package arrays_hashing

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"alpha", args{[]string{"neet", "code", "love", "you"}}, "4-neet4-code4-love3-you"},
		{"ascii", args{[]string{"we", "say", ":", "yes"}}, "2-we3-say1-:3-yes"},
		{"long", args{[]string{"hallelulah", "girl"}}, "10-hallelulah4-girl"},
		{"long-multi", args{[]string{"hallelulahh", "giiiiiirllll"}}, "11-hallelulahh12-giiiiiirllll"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.strs); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"alpha", args{"4-neet4-code4-love3-you"}, []string{"neet", "code", "love", "you"}},
		{"ascii", args{"2-we3-say1-:3-yes"}, []string{"we", "say", ":", "yes"}},
		{"long", args{"10-hallelulah4-girl"}, []string{"hallelulah", "girl"}},
		{"long-multi", args{"11-hallelulahh12-giiiiiirllll"}, []string{"hallelulahh", "giiiiiirllll"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
