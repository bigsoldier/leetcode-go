/**
 *Created by Xie Jian on 2020/1/16 10:48
 */
package _009_Palindrome_Number

//Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
// Example 1:
//
//
//Input: 121
//Output: true
//
//
// Example 2:
//
//
//Input: -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//
//
// Example 3:
//
//
//Input: 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
//
// Follow up:
//
// Coud you solve it without converting the integer to a string?
// Related Topics Math

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var temp int
	var y = x
	for y != 0 {
		temp = temp*10 + y%10
		y /= 10
	}
	if x == temp {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
