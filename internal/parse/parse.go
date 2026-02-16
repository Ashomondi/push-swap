package parse

import (
	"errors"
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	if len(args) == 0 {
		return []int{}, nil
	}

	var nums []int
	seen := make(map[int]bool)

	for _, arg := range args {
		parts := strings.Fields(arg) // handles multiple spaces nicely
		for _, p := range parts {
			n64, err := strconv.ParseInt(p, 10, 32)
			if err != nil {
				return nil, errors.New("invalid integer")
			}
			n := int(n64)

			if seen[n] {
				return nil, errors.New("duplicate")
			}
			seen[n] = true
			nums = append(nums, n)
		}
	}

	return nums, nil
}
