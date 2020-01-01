package p90

func Invoke(s string) int {
	i := len(s)
	for {
		flag := true
		for j := 0; j < len(s); j++ {
			if i-j-1 < len(s) && s[j] != s[i-j-1] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
		i++
	}
}
