package p98

var grid = make([][]bool, 100)
var vx = []int{1, -1, 0, 0}
var vy = []int{0, 0, 1, -1}
var prob = make([]float64, 4)

func dfs(x int, y int, n int) float64 {
	if grid[x][y] {
		return 0
	}
	if n == 0 {
		return 1
	}

	grid[x][y] = true
	var ret float64 = 0
	for i, _ := range prob {
		ret += dfs(x+vx[i], y+vy[i], n-1) * prob[i]
	}
	grid[x][y] = false

	return ret
}

func Invoke(n int, east int, west int, south int, north int) float64 {
	for i, _ := range grid {
		grid[i] = make([]bool, 100)
	}

	prob[0] = float64(east) / 100.0
	prob[1] = float64(west) / 100.0
	prob[2] = float64(south) / 100.0
	prob[3] = float64(north) / 100.0

	return dfs(50, 50, n)
}
