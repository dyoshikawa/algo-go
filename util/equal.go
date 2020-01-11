package util

func EqualArrInt(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EqualDoubleArrInt(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if !EqualArrInt(a[i], b[i]) {
			return false
		}
	}
	return true
}
