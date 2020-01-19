package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&m)
	okqs := make([]bool, n)
	ngqs := make([]int, n)
	for i := 0; i < m; i++ {
		var p int
		_, _ = fmt.Scan(&p)
		p--
		var s string
		_, _ = fmt.Scan(&s)
		if !okqs[p] {
			if s == "AC" {
				okqs[p] = true
			} else {
				ngqs[p]++
			}
		}
	}
	ok := 0
	ng := 0
	for i, v := range okqs {
		if v {
			ok++
			ng += ngqs[i]
		}
	}
	res := fmt.Sprintf("%d %d", ok, ng)
	fmt.Println(res)
}
