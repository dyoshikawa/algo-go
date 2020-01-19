package a

import "fmt"

func main() {
	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var s string
	_, _ = fmt.Scan(&s)
	var res int
	for i, v := range strs {
		if s == v {
			res = i
		}
	}
	fmt.Println(strs[res+1])
}
