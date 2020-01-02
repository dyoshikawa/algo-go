package p119

func search(arr []int, n int, key int) int {
	i := 0
	// 番兵
	arr[n] = key
	for arr[i] != key {
		i++
	}
	if i == n {
		return -1
	}
	return i
}

func Invoke(n int, S []int, q int, T []int) int {
	// 番兵用要素を一つ足す
	S = append(S, -1)
	res := 0
	for _, v := range T {
		if search(S, n, v) != -1 {
			res++
		}
	}

	return res
}
