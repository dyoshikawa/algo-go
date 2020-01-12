package collection_int

func EqualArr(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EqualDoubleArr(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if !EqualArr(a[i], b[i]) {
			return false
		}
	}
	return true
}

// 参考
// https://qiita.com/usk81/items/5ff7bfe27f702d77e909
func UnsetArr(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

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
	last := len(s.Data) - 1
	popped := s.Data[last]
	s.Data = UnsetArr(s.Data, last)
	return popped
}

func (s *Stack) Exists() bool {
	return len(s.Data) > 0
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}
