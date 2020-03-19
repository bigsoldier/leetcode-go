package code

import "testing"

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := 6
	result := trap(height)
	t.Log("result is ", result, " ", result == ret)
	if result != ret {
		t.Fatal("this question is not solve")
	}
}
