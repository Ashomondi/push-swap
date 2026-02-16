package algo

import "push-swap/internal/ops"

func RadixSort(ex *ops.Executor) {
	n := ex.A.Len()
	if n <= 1 {
		return
	}

	// Find max bits needed
	max := 0
	for _, v := range ex.A.Data {
		if v > max {
			max = v
		}
	}

	maxBits := 0
	for (max >> maxBits) != 0 {
		maxBits++
	}

	for bit := 0; bit < maxBits; bit++ {
		for i := 0; i < n; i++ {
			top := ex.A.Data[0]
			if ((top >> bit) & 1) == 0 {
				ex.Pb()
			} else {
				ex.Ra()
			}
		}
		for ex.B.Len() > 0 {
			ex.Pa()
		}
	}
}
