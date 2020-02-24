/**
 *Created by Xie Jian on 2020/1/10 9:40
 */
package _006_ZigZag_Conversion

import "strings"

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//
//
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
//
//string convert(string s, int numRows);
//
// Example 1:
//
//
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
//
//
// Example 2:
//
//
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
// Related Topics String

//leetcode submit region begin(Prohibit modification and deletion)
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

//leetcode submit region end(Prohibit modification and deletion)
