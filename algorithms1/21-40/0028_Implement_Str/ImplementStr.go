/**
 *Created by Xie Jian on 2020/2/18 16:36
 */
package _028_Implement_Str

//Implement strStr().
//
// Return the index of the first occurrence of needle in haystack, or -1 if need
//le is not part of haystack.
//
// Example 1:
//
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//
//
// Example 2:
//
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//
//
// Clarification:
//
// What should we return when needle is an empty string? This is a great questio
//n to ask during an interview.
//
// For the purpose of this problem, we will return 0 when needle is an empty str
//ing. This is consistent to C's strstr() and Java's indexOf().
// Related Topics Two Pointers String

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		var j int
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
