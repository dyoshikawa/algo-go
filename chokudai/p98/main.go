package p98

func Invoke(friends []string) int {
	ans := 0
	n := len(friends)

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if friends[i][j] == "Y" {

			} else {

			}
		}
	}

	return 0
}
