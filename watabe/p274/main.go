package p269

var visited []bool
var graph map[int][]int
var cnt int
var res [][]int

//func InvokeStack(n int, as [][]int) [][]int {
//	graph = make(map[int][]int)
//	for i, a := range as {
//		for j, v := range a {
//			if j == 0 || j == 1 {
//				continue
//			}
//			graph[i] = append(graph[i], v-1)
//		}
//	}
//
//	visited = make([]bool, n)
//	cnt = 0
//	res = make([][]int, n)
//	for i, _ := range res {
//		res[i] = make([]int, 2)
//	}
//
//	stk := collection_int.NewStack()
//	now := 0
//	stk.Push(now)
//	for stk.Exists() {
//		visited[now] = true
//		cnt++
//		res[now][0] = cnt
//
//		for _, v := range graph[now] {
//			if !visited[v] {
//				stk.Push(v)
//			}
//		}
//
//		now = stk.Pop()
//	}
//
//	return res
//}

func dfs(a int) {
	visited[a] = true
	cnt++
	res[a][0] = cnt

	for _, v := range graph[a] {
		if !visited[v] {
			dfs(v)
		}
	}

	cnt++
	res[a][1] = cnt
}

func InvokeRecursion(n int, as [][]int) [][]int {
	graph = make(map[int][]int)
	for i, a := range as {
		for j, v := range a {
			if j == 0 || j == 1 {
				continue
			}
			graph[i] = append(graph[i], v-1)
		}
	}

	visited = make([]bool, n)
	cnt = 0
	res = make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, 2)
	}
	dfs(0)

	return res
}
