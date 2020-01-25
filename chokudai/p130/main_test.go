package p130

import "testing"

func TestInvoke0(t *testing.T) {
	maze := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}
	startRow := 0
	startCol := 1
	moveRow := []int{1, 0, -1, 0}
	moveCol := []int{0, 1, 0, -1}
	if res := Invoke(maze, startRow, startCol, moveRow, moveCol); res != 3 {
		t.Fatal(res)
	}
}

func TestInvoke1(t *testing.T) {
	maze := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}
	startRow := 0
	startCol := 1
	moveRow := []int{1, 0, -1, 0, 1, 1, -1, -1}
	moveCol := []int{0, 1, 0, -1, 1, -1, 1, -1}
	if res := Invoke(maze, startRow, startCol, moveRow, moveCol); res != 2 {
		t.Fatal(res)
	}
}

func TestInvoke2(t *testing.T) {
	maze := [][]string{
		{"x", ".", "x"},
		{"x", "x", "x"},
		{"x", ".", "x"},
	}
	startRow := 0
	startCol := 1
	moveRow := []int{1, 0, -1, 0}
	moveCol := []int{0, 1, 0, -1}
	if res := Invoke(maze, startRow, startCol, moveRow, moveCol); res != -1 {
		t.Fatal(res)
	}
}
