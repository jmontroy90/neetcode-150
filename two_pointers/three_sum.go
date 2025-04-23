package two_pointers

import "slices"

func ThreeSum(numbers []int) [][]int {
	slices.Sort(numbers)
	var trips [][]int
	for i := 0; i < len(numbers); i++ {
		j, k := i+1, len(numbers)-1
		for j < k {
			if target := numbers[i] * -1; target < numbers[j]+numbers[k] {
				k--
			} else if target > numbers[j]+numbers[k] {
				j++
			} else {
				trips = append(trips, []int{numbers[i], numbers[j], numbers[k]})
				break
			}
		}
	}
	return trips
}
