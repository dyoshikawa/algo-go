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
	// 参考
	// https://qiita.com/egnr-in-6matroom/items/282aa2fd117aab9469bd#slice%E3%81%AE%E5%85%88%E9%A0%AD%E3%81%AB%E8%A6%81%E7%B4%A0%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B
	s.Data, s.Data[0] = append(s.Data[0:1], s.Data[0:]...), a
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
