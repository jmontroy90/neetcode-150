package arrays_hashing

import (
	"slices"
	"unicode/utf8"
)

// utf8-safe
func IsAnagram(s, t string) bool {
	sr := []rune(s)
	tr := []rune(t)
	slices.Sort(sr)
	slices.Sort(tr)
	for i := range sr {
		if sr[i] != tr[i] {
			return false
		}
	}
	return true
}

// utf8-safe
func IsAnagramFreq(s, t string) bool {
	slr := utf8.RuneCountInString(s)
	tlr := utf8.RuneCountInString(t)
	if slr != tlr {
		return false
	}
	sf := make(map[rune]int, slr)
	tf := make(map[rune]int, tlr)
	for i := 0; i < slr; i++ {
		// be careful how you slice strings, deal with runes as much as possible.
		sf[[]rune(s)[i]]++
		tf[[]rune(t)[i]]++
	}
	if len(sf) != len(tf) {
		return false // probably not much of a time save
	}
	for r, c := range sf {
		if tf[r] != c {
			return false
		}
	}
	return true
}
