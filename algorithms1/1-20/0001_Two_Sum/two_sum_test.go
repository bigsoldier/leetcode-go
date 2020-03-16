/**
 *Created by Xie Jian on 2019/9/16 14:47
 */
package leetcode

import (
	"fmt"
	"github.com/betterfor/leetcode-go/algorithms1/1-20"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4},
		{2, 7, 11, 5},
		{4, 3, 4, 5},
		{6, 4, 3, 1},
	}
	targets := []int{5, 9, 8, 7}

	results := [][]int{
		{1, 2},
		{0, 1},
		{0, 2},
		{1, 2},
	}
	fmt.Println("-------------------Leetcode 1-----------------------")
	for i := 0; i < len(tests); i++ {
		ret := __20.twoSum(tests[i], targets[i])
		fmt.Printf("num=%v target=%v result=%v\n", tests[i], targets[i], ret)
		if ret[1] != results[i][1] {
			t.Fatalf("case %d fail:%v\n", i, ret)
		}
	}
}
