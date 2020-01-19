package main

import "fmt"

func main() {
	var n, k, m int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&k)
	_, _ = fmt.Scan(&m)
	as := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		var a int
		_, _ = fmt.Scan(&a)
		as[i] = a
	}
	expectedTotal := m * n
	actualTotal := 0
	for _, v := range as {
		actualTotal += v
	}
	res := expectedTotal - actualTotal
	if res > k {
		fmt.Println(-1)
	} else if res <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(res)
	}
}
