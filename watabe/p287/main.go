package p287

var n, m, q int
var fs [][]int
var qs [][]int
var visited []bool
var graph map[int][]int
var goal int
var ok bool

func dfs(a int) {
	visited[a] = true
	if a == goal {
		ok = true
	}

	for _, v := range graph[a] {
		if !visited[v] {
			dfs(v)
		}
	}
}

func Invoke() []bool {
	visited = make([]bool, n)
	graph = make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}
	for _, f := range fs {
		graph[f[0]] = append(graph[f[0]], f[1])
		graph[f[1]] = append(graph[f[1]], f[0])
	}

	res := make([]bool, q)
	for i, v := range qs {
		ok = false
		goal = v[1]
		dfs(v[0])
		res[i] = ok
	}

	return res
}
