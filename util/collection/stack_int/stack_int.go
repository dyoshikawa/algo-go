package stack_int

import "github.com/dyoshikawa/algo/util/collection/util_int"

type Stack struct {
	Data []int
}

func NewStack() *Stack {
	return &Stack{
		Data: make([]int, 0),
	}
}

func (s *Stack) Push(a int) {
	s.Data = append([]int{a}, s.Data...)
}

func (s *Stack) Pop() int {
	popped := s.Data[0]
	s.Data = util_int.UnsetArr(s.Data, 0)
	return popped
}

func (s *Stack) Exists() bool {
	return len(s.Data) > 0
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}
