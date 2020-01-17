package util_string

func EqualArr(a []string, b []string) bool {
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

func EqualDoubleArr(a [][]string, b [][]string) bool {
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
