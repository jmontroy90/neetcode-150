package arrays_hashing

// Original idea: sparse array
func LongestConsecutive(nums []int) int {
	arr := make([]bool, 100)
	for _, n := range nums {
		if n > len(arr) {
			arr = resizeSlice(arr, n)
		}
		arr[n] = true
	}
	var longestStreak, currStreak int
	if arr[0] {
		currStreak = 1 // OOB edge case
	}
	// this logic is kinda isomorphic to the optimal solution, I think
	for ix := 1; ix < len(arr); ix++ {
		if !arr[ix-1] && arr[ix] {
			currStreak = 1
		} else if arr[ix] {
			currStreak++
		}
		if currStreak > longestStreak {
			longestStreak = currStreak
		}
	}
	return longestStreak
}

func resizeSlice(old []bool, n int) []bool {
	newArr := make([]bool, n)
	copy(newArr, old)
	return newArr
}

// Optimal solution: hash set
func LongestConsecutiveHash(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		set[n] = true
	}
	var longest, curr int // streaks
	for _, n := range nums {
		currNum := n
		if isSet := set[currNum-1]; !isSet {
			isSet = true
			for isSet {
				isSet = set[currNum+1]
				curr++
				currNum++
			}
			if curr > longest {
				longest = curr
			}
		}
		curr = 0
	}
	return longest
}
