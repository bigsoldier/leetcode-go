/**
 *Created by Xie Jian on 2020/1/15 10:08
 */
package _007_Reverse_Integer

import "math"

//Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
//
//
//Input: 123
//Output: 321
//
//
// Example 2:
//
//
//Input: -123
//Output: -321
//
//
// Example 3:
//
//
//Input: 120
//Output: 21
//
//
// Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
// Related Topics Math

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var temp int
	for ; x != 0; x /= 10 {
		t := x % 10
		if temp > math.MaxInt32/10 || (temp == math.MaxInt32/10 && t > 7) {
			return 0
		}
		if temp < math.MinInt32/10 || (temp == math.MinInt32/10 && t < -8) {
			return 0
		}
		temp = temp*10 + t
	}
	return temp
}

//leetcode submit region end(Prohibit modification and deletion)
