/**
 *Created by Xie Jian on 2020/2/14 12:58
 */
package _020_Valid_Parentheses

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
//
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
//
//
// Note that an empty string is also considered valid.
//
// Example 1:
//
//
//Input: "()"
//Output: true
//
//
// Example 2:
//
//
//Input: "()[]{}"
//Output: true
//
//
// Example 3:
//
//
//Input: "(]"
//Output: false
//
//
// Example 4:
//
//
//Input: "([)]"
//Output: false
//
//
// Example 5:
//
//
//Input: "{[]}"
//Output: true
//
// Related Topics String Stack

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' ||
			v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' ||
			v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
