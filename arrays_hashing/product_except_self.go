package arrays_hashing

import "slices"

// This works, but degenerates to O(n^2) if every element is zeroes.
// You could add a check for how many zeroes, because if there's more than one, every product will be zero.
func ProductExceptSelfDivision(nums []int) []int {
	product := sliceProduct(nums)
	productExceptSelf := make([]int, len(nums))
	for ix := range nums {
		divisor := nums[ix]
		if divisor == 0 {
			productExceptSelf[ix] = sliceProduct(slices.Concat(nums[0:ix], nums[ix+1:]))
			continue
		}
		productExceptSelf[ix] = product / divisor
	}
	return productExceptSelf
}

func sliceProduct(nums []int) int {
	product := nums[0]
	for _, num := range nums[1:] {
		product *= num
	}
	return product
}

func ProductExceptSelfPrefixSuffix(nums []int) []int {
	prefix, suffix := make([]int, len(nums)+1), make([]int, len(nums)+1)
	prefix[0], suffix[len(nums)] = 1, 1 // this essentially gives the zero value to this pseudo-fold left / fold right
	for ix := 1; ix <= len(nums); ix++ {
		prefix[ix] = prefix[ix-1] * nums[ix-1]
		suffix[len(nums)-ix] = suffix[len(nums)-ix+1] * nums[len(nums)-ix]
	}
	out := make([]int, len(nums))
	for ix := 0; ix < len(nums); ix++ {
		out[ix] = prefix[ix] * suffix[ix+1]
	}
	return out
}
