package arrays_hashing

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"anagram", args{"racecar", "carrace"}, true},
		{"anagram-utf8", args{"racécar", "carracé"}, true},
		{"not-anagram", args{"jar", "jam"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnagramFreq(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"anagram", args{"racecar", "carrace"}, true},
		{"anagram-utf8", args{"racécar", "carracé"}, true},
		{"not-anagram", args{"jar", "jam"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagramFreq(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagramFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}
