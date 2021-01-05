package code

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) (ans string) {
	var strs []string
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.SliceStable(strs, func(i, j int) bool {
		a, b := strs[i], strs[j]
		return a+b > b+a
	})
	if strs[0] == "0" {
		return "0"
	}
	for _, str := range strs {
		ans += str
	}
	return
}
