package arrays_hashing

import (
	"slices"
)

// Brute force: Check each number against all others, O(n^2)
// Sorting: pdqsort (pattern-defeating quicksort) is O(nk), where k == num distinct	elements
func HasDuplicate(nums []int) bool {
	// This is unnecessary, since []int is ordered
	slices.SortFunc(nums, func(a, b int) int {
		if a < b {
			return -1
		} else if a == b {
			return 0
		} else {
			return 1
		}
	})
	for ix := 0; ix < len(nums)-1; ix++ {
		if nums[ix] == nums[ix+1] {
			return true
		}
	}
	return false
}

// Hash set: O(n) time and space
func HasDuplicateHash(nums []int) bool {
	nm := make(map[int]bool, len(nums))
	for _, n := range nums {
		if nm[n] {
			return true
		}
		nm[n] = true
	}
	return false
}
