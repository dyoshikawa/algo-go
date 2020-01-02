package p136

var n, k int
var T []int

func check(P int) int {
	i := 0
	for j := 0; j < k; j++ {
		s := 0
		for s+T[i] <= P {
			s += T[i]
			i++
			if i == n {
				return n
			}
		}
	}
	return i
}

func solve() int {
	left := 0
	right := 100000 * 10000
	var mid int
	for right-left > 1 {
		mid = (left + right) / 2
		v := check(mid)
		if v >= n {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

func Invoke(argN int, argK int, argT []int) int {
	n = argN
	k = argK
	T = argT
	return solve()
}
