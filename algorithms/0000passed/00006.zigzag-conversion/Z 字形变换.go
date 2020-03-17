package code

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret = make([]string, numRows)
	var T = 2*numRows - 2 // 周期（必定为偶数）
	for index, bt := range []byte(s) {
		rows := index % T
		if rows > T/2 {
			ret[T-rows] += string(bt)
		} else {
			ret[rows] += string(bt)
		}
	}
	return strings.Join(ret, "")
}
