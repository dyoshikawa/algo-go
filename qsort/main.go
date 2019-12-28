package main

import "fmt"

func qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pv := arr[0]
		var less []int
		for i := 1; i < len(arr); i++ {
			if arr[i] <= pv {
				less = append(less, arr[i])
			}
		}
		var grtr []int
		for i := 1; i < len(arr); i++ {
			if arr[i] > pv {
				grtr = append(grtr, arr[i])
			}
		}
		newArr := []int{pv}
		newArr = append(qsort(less), newArr...)
		newArr = append(newArr, qsort(grtr)...)
		return newArr
	}
}

func main() {
	fmt.Println(qsort([]int{10, 5, 2, 3}))
}