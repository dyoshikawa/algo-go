package p240

import (
	"github.com/dyoshikawa/algo/util/collection/util_int"
	"testing"
)

func TestInvoke(t *testing.T) {
	as := []string{
		"insert 8",
		"insert 2",
		"extract",
		"insert 10",
		"extract",
		"insert 11",
		"extract",
		"extract",
		"end",
	}
	expected := []int{
		8, 10, 11, 2,
	}
	if res := Invoke(as); !util_int.EqualArr(res, expected) {
		t.Fatal(res)
	}
}

func TestInvokeAlt(t *testing.T) {
	as := []string{
		"insert 8",
		"insert 2",
		"extract",
		"insert 10",
		"extract",
		"insert 11",
		"extract",
		"extract",
		"end",
	}
	expected := []int{
		8, 10, 11, 2,
	}
	if res := InvokeAlt(as); !util_int.EqualArr(res, expected) {
		t.Fatal(res)
	}
}
