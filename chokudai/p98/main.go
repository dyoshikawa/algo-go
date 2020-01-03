package p98

func Invoke(friends [][]string) int {
	ans := 0
	n := len(friends)

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if friends[i][j] == "Y" {
				cnt++
			} else {
				for k := 0; k < n; k++ {
					if friends[j][k] == "Y" && friends[k][i] == "Y" {
						cnt++
						break
					}
				}
			}
		}
		if ans < cnt {
			ans = cnt
		}
	}

	return ans
}
