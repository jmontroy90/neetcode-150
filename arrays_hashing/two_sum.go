package arrays_hashing

// brute force: O(n^2)
func TwoSum(nums []int, target int) []int {
	for i1 := 0; i1 < len(nums); i1++ {
		for i2 := i1 + 1; i2 < len(nums); i2++ {
			if nums[i1]+nums[i2] == target {
				return []int{i1, i2}
			}
		}
	}
	return nil
}

// O(n) - looking for difference is neat, not sure if I thought of it myself or just saw it a long long time ago
func TwoSumHash(nums []int, target int) []int {
	h := make(map[int]int, len(nums))
	for ix, n := range nums {
		if ix2, ok := h[target-n]; ok {
			return []int{ix2, ix}
		}
		h[n] = ix
	}
	return nil
}
