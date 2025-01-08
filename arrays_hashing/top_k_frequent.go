package arrays_hashing

func TopKFrequent(nums []int, k int) []int {
	if k == 0 {
		return nil // lil paranoid edge case
	}
	fs := make(map[int]int, len(nums))
	for _, num := range nums {
		fs[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for n, c := range fs {
		// This is crafty - the index is our count, and counts are ordered.
		buckets[c] = append(buckets[c], n)
	}
	var out []int
	for i := len(buckets) - 1; i > 0; i-- {
		for _, n := range buckets[i] {
			out = append(out, n)
			if len(out) == k {
				return out
			}
		}
	}
	return nil
}
