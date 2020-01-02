package p124

func binarySearch(arr []int, n int, key int) int {
	left := 0
	right := n
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == key {
			return mid
		} else if key < arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Invoke(n int, S []int, q int, T []int) int {
	res := 0
	for _, v := range T {
		if binarySearch(S, n, v) != -1 {
			res++
		}
	}

	return res
}
