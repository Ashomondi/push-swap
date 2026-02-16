package algo

import "push-swap/internal/stack"

func IsSortedAsc(s *stack.Stack) bool {
	if s.Len() < 2 {
		return true
	}
	for i := 0; i < s.Len()-1; i++ {
		if s.Data[i] > s.Data[i+1] {
			return false
		}
	}
	return true
}
