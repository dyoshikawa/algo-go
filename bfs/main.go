package main

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{data: make([]int, 0)}
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() int {
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

var q *Queue
var g map[int][]int
var done []bool
var start int
var goal int
var cnt int

func bfs(i int) bool {
	if done[i] {
		return bfs(q.Pop())
	}
	done[i] = true
	cnt++
	if i == goal {
		return true
	}
	for _, v := range g[i] {
		q.Push(v)
	}
	if q.IsEmpty() {
		return false
	} else {
		return bfs(q.Pop())
	}
}

func main() {
	q = NewQueue()
	g = make(map[int][]int)
	g[0] = []int{}
	g[1] = []int{0}
	g[2] = []int{1}
	g[3] = []int{2}
	done = make([]bool, 4)
	start = 3
	goal = 1
	for _, v := range g[start] {
		q.Push(v)
	}
	done[start] = true
	bfs(start)
	fmt.Println(cnt)
}
