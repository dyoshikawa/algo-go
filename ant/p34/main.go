package p34

var n int
var a []int
var k int

func dfs(i int, sum int) bool {
	if n == i {
		return sum == k
	}

	if dfs(i+1, sum) {
		return true
	}

	if dfs(i+1, sum+a[i]) {
		return true
	}

	return false
}

func Invoke(argN int, argA []int, argK int) string {
	n = argN
	a = argA
	k = argK

	if dfs(0, 0) {
		return "Yes"
	}
	return "No"
}
