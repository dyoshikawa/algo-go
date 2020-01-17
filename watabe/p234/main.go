package p234

import "fmt"

func Invoke(n int, as []int) []string {
	res := make([]string, n)
	for i, _ := range as {
		i++
		s := fmt.Sprintf("node %d: ", i)
		key := as[i-1]
		s += fmt.Sprintf("key = %d, ", key)
		if pi := i / 2; pi >= 1 {
			parentKey := as[pi-1]
			s += fmt.Sprintf("parent key = %d, ", parentKey)
		}
		if li := i * 2; li <= n {
			leftKey := as[li-1]
			s += fmt.Sprintf("left key = %d, ", leftKey)
		}
		if ri := i*2 + 1; ri <= n {
			rightKey := as[ri-1]
			s += fmt.Sprintf("right key = %d, ", rightKey)
		}
		res[i-1] = s
	}
	return res
}
