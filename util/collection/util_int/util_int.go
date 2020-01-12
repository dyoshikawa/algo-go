package util_int

func EqualArr(a []int, b []int) bool {
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

func EqualDoubleArr(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if !EqualArr(a[i], b[i]) {
			return false
		}
	}
	return true
}

// 参考
// https://qiita.com/usk81/items/5ff7bfe27f702d77e909
func UnsetArr(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
