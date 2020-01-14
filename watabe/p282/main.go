package p269

import (
	"github.com/dyoshikawa/algo/util/collection/queue_int"
)

func Invoke(n int, as [][]int) [][]int {
	graph := make(map[int][]int)
	for i, a := range as {
		for j, v := range a {
			if j == 0 || j == 1 {
				continue
			}
			graph[i] = append(graph[i], v-1)
		}
	}

	visited := make([]bool, n)
	start := 0
	ds := make([]int, n)
	for i, _ := range ds {
		ds[i] = -1
	}
	ds[start] = 0
	q := queue_int.NewQueue()
	q.Push(start)
	for q.Exists() {
		now := q.Pop()
		visited[now] = true

		for _, v := range graph[now] {
			if !visited[v] {
				q.Push(v)
				ds[v] = ds[now] + 1
			}
		}
	}

	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, 2)
	}
	for i, _ := range ds {
		res[i][0] = i + 1
		res[i][1] = ds[i]
	}
	return res
}
