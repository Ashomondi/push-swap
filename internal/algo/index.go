package algo

import (
	"sort"
)

func ToIndexes(nums []int) []int {
	cp := make([]int, len(nums))
	copy(cp, nums)

	sort.Ints(cp)

	// map value -> index
	pos := make(map[int]int, len(cp))
	for i, v := range cp {
		pos[v] = i
	}

	out := make([]int, len(nums))
	for i, v := range nums {
		out[i] = pos[v]
	}
	return out
}
