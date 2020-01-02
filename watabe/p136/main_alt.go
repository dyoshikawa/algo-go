package p136

func InvokeAlt(n int, k int, w []int) int {
	sum := 0
	for _, v := range w {
		sum += v
	}
	avg := sum / k
	if sum%k > 0 {
		// 切り上げ
		avg++
	}

	max := 0
	for _, v := range w {
		if max < v {
			max = v
		}
	}

	// 大きい方を返す
	if avg < max {
		return max
	} else {
		return avg
	}
}
