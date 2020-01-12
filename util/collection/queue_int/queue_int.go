package queue_int

import "github.com/dyoshikawa/algo/util/collection/util_int"

type Queue struct {
	Data []int
}

func NewQueue() *Queue {
	return &Queue{
		Data: make([]int, 0),
	}
}

func (s *Queue) Push(a int) {
	s.Data = append(s.Data, a)
}

func (s *Queue) Pop() int {
	last := len(s.Data) - 1
	popped := s.Data[last]
	s.Data = util_int.UnsetArr(s.Data, last)
	return popped
}

func (s *Queue) Exists() bool {
	return len(s.Data) > 0
}

func (s *Queue) IsEmpty() bool {
	return len(s.Data) == 0
}
