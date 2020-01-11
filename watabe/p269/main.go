package p269

import "fmt"

func Invoke(n int, as [][]int) [][]int {
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}

	for i, a := range as {
		for j, v := range a {
			if j == 0 || j == 1 {
				continue
			}
			res[i][v-1] = 1
		}
	}

	fmt.Println(res)
	return res
}
