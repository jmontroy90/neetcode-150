package arrays_hashing

// This uses as the key just the sum of all runes as integers per word.
// I prevent strs of different lengths being equal by padding out the sum with the length * 10000.
// There's a lot of talk of [26]rune, but that's obviously no good without ASCII.
// Mine does work without ASCII.
func GroupAnagrams(strs []string) [][]string {
	mm := make(map[int][]string, len(strs))
	for _, str := range strs {
		var k int
		for _, r := range []rune(str) {
			k += int(r)
		}
		k += len(str) * 10000 // this confounds false positives, i think
		mm[k] = append(mm[k], str)
	}
	var out [][]string
	for _, ga := range mm {
		out = append(out, ga)
	}
	return out
}
