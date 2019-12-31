package p35

var n int
var m int
var pic [][]string

func dfs(x int, y int) {
	pic[x][y] = "."

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx := x + dx
			ny := y + dy

			if 0 <= nx && nx < n && 0 <= ny && ny < m && pic[nx][ny] == "W" {
				dfs(nx, ny)
			}
		}
	}

}

func Invoke(argN int, argM int, argPic [][]string) int {
	n = argN
	m = argM
	pic = argPic

	res := 0
	for i, v := range pic {
		for j, w := range v {
			if w == "W" {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
