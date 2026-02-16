package algo

import "push-swap/internal/ops"

func Sort2(ex *ops.Executor) {
	if ex.A.Len() != 2 {
		return
	}
	if ex.A.Data[0] > ex.A.Data[1] {
		ex.Sa()
	}
}

func Sort3(ex *ops.Executor) {
	if ex.A.Len() != 3 {
		return
	}

	a := ex.A.Data[0]
	b := ex.A.Data[1]
	c := ex.A.Data[2]

	// Already sorted
	if a < b && b < c {
		return
	}

	// Case: 2 1 3
	if a > b && b < c && a < c {
		ex.Sa()
		return
	}

	// Case: 3 2 1
	if a > b && b > c {
		ex.Sa()
		ex.Rra()
		return
	}

	// Case: 3 1 2
	if a > b && b < c && a > c {
		ex.Ra()
		return
	}

	// Case: 1 3 2
	if a < b && b > c && a < c {
		ex.Sa()
		ex.Ra()
		return
	}

	// Case: 2 3 1
	if a < b && b > c && a > c {
		ex.Rra()
		return
	}
}
func minIndex(a []int) int {
	minI := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[minI] {
			minI = i
		}
	}
	return minI
}

func bringIndexToTopA(ex *ops.Executor, idx int) {
	n := ex.A.Len()
	if n == 0 {
		return
	}

	// If idx is in first half -> rotate up
	if idx <= n/2 {
		for idx > 0 {
			ex.Ra()
			idx--
		}
	} else {
		// rotate down
		steps := n - idx
		for steps > 0 {
			ex.Rra()
			steps--
		}
	}
}

func Sort4or5(ex *ops.Executor) {
	n := ex.A.Len()
	if n != 4 && n != 5 {
		return
	}

	// push smallest until 3 left in A
	for ex.A.Len() > 3 {
		idx := minIndex(ex.A.Data)
		bringIndexToTopA(ex, idx)
		ex.Pb()
	}

	Sort3(ex)

	// push back
	for ex.B.Len() > 0 {
		ex.Pa()
	}
}
func SortSmall(ex *ops.Executor) {
	n := ex.A.Len()
	if n <= 5 {
		return
	}

	// Push smallest values to B until 3 left in A
	for ex.A.Len() > 3 {
		idx := minIndex(ex.A.Data)
		bringIndexToTopA(ex, idx)
		ex.Pb()
	}

	// Sort remaining 3 in A
	Sort3(ex)

	// Push back everything from B to A in correct order
	for ex.B.Len() > 0 {
		ex.Pa()
	}

	// Final rotate: bring smallest to top
	idx := minIndex(ex.A.Data)
	bringIndexToTopA(ex, idx)
}

