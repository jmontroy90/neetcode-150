package two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{"Was it a car or a cat I saw?"}, true},
		{"true-with-nums", args{"W1as 2it a car or a cat I 2sa1w?"}, true},
		{"false", args{"tab a cat"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
