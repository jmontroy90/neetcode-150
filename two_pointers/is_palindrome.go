package two_pointers

// definitely possible to do this in place without a cleaned version too - you just need to distinctly use two pointers and iterate each past any character you don't want to evaluate.
// other neetcode solution: reverse strings and compare
func IsPalindrome(s string) bool {
	reduced := make([]rune, len(s))
	var rIx int
	for _, c := range s {
		if 'A' <= c && c <= 'Z' { // uppercase
			reduced[rIx] = c + ('a' - 'A')
			rIx++
		} else if 'a' <= c && c <= 'z' || '0' <= c && c <= '9' { // lowercase
			reduced[rIx] = c
			rIx++
		}
	}
	clean := string(reduced[:rIx])
	for lp := range clean {
		rp := len(clean) - 1 - lp
		if lp > rp {
			break
		}
		if clean[lp] != clean[len(clean)-1-lp] {
			return false
		}
	}
	return true
}
