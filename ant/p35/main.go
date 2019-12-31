package p35

var n int
var m int
var field [][]string

func dfs(x int, y int) {
	field[x][y] = "."

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx := x + dx
			ny := y + dy

			if 0 <= nx && nx < n && 0 <= ny && ny < m && field[nx][ny] == "W" {
				dfs(nx, ny)
			}
		}
	}

}

func Invoke(argN int, argM int, argField [][]string) int {
	n = argN
	m = argM
	field = argField

	res := 0
	for i, v := range field {
		for j, w := range v {
			if w == "W" {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
