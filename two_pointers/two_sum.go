package two_pointers

// this is like a junky binary search, as in it roughly guesses where the next number might be, and then scans
// if the numbers and their indexes are completely unrelated, this degenerates hard. not an ideal solution
func TwoSum(numbers []int, target int) []int {
	for lp := range numbers {
		rp := target - numbers[lp]
		if rp >= len(numbers) {
			rp = len(numbers) - 1
		}
		guess := numbers[lp] + numbers[rp]
		for {
			if guess == target {
				return []int{lp, rp}
			} else if guess > target {
				rp--
			} else {
				rp++
			}
			guess = numbers[lp] + numbers[rp]
		}
	}
	return nil
}

func TwoSumPointers(numbers []int, target int) []int {
	lp, rp := 0, len(numbers)-1
	guess := numbers[lp] + numbers[rp]
	for {
		if guess > target {
			rp--
		} else if guess < target {
			lp++
		} else {
			return []int{lp, rp}
		}
		guess = numbers[lp] + numbers[rp]
	}
}
