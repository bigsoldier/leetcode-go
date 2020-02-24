/**
 *Created by Xie Jian on 2020/1/16 10:49
 */
package _010_Regular_Expression_Matching

//Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
//
//
//'.' Matches any single character.
//'*' Matches zero or more of the preceding element.
//
//
// The matching should cover the entire input string (not partial).
//
// Note:
//
//
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like . or *.
//
//
// Example 1:
//
//
//Input:
//s = "aa"
//p = "a"
//Output: false
//Explanation: "a" does not match the entire string "aa".
//
//
// Example 2:
//
//
//Input:
//s = "aa"
//p = "a*"
//Output: true
//Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
//
//
// Example 3:
//
//
//Input:
//s = "ab"
//p = ".*"
//Output: true
//Explanation: ".*" means "zero or more (*) of any character (.)".
//
//
// Example 4:
//
//
//Input:
//s = "aab"
//p = "c*a*b"
//Output: true
//Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
//
//
// Example 5:
//
//
//Input:
//s = "mississippi"
//p = "mis*is*p*."
//Output: false
//
// Related Topics String Dynamic Programming Backtracking

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配 */

	dp[0][0] = true
	/**
	 * 根据题目的设定， "" 可以与 "a*b*c*" 相匹配
	 * 但""不与 "a*b*c*d" 相匹配
	 * 所以，需要把奇数位上有 "*" 的 dp 设置成 true
	 */
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				/* p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上 */
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				/* 此时，p[j] 的匹配情况与 p[j-1] 的内容相关。 */
				if p[j-1] != s[i] && p[j-1] != '.' {
					/**
					 * p[j] 无法与 s[i] 匹配上
					 * p[j-1:j+1] 只能被当做 ""
					 */
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/**
					 * p[j] 与 s[i] 匹配上
					 * p[j-1;j+1] 作为 "x*", 可以有三种解释
					 */
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" 解释为 "" */
						dp[i+1][j] || /* "x*" 解释为 "x" */
						dp[i][j+1] /* "x*" 解释为 "xx..." */
				}
			}
		}
	}

	return dp[sSize][pSize]
}

//leetcode submit region end(Prohibit modification and deletion)
