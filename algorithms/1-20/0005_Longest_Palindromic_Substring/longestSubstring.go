/**
 *Created by Xie Jian on 2019/12/18 15:20
 */
package _005_Longest_Palindromic_Substring

//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
// Example 1:
//
//
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//
//
// Example 2:
//
//
//Input: "cbbd"
//Output: "bb"
//
// Related Topics String Dynamic Programming

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// P[i][j]表示回文，P[i+1][j-1]也是回文
	longest := s[0:1]
	for i := 1; i < len(s); i++ {
		for rightStep := 0; rightStep < 2; rightStep++ {
			for p, q := i-1, i+rightStep; p >= 0 && q < len(s) && s[p] == s[q]; {
				if q-p+1 > len(longest) {
					longest = s[p : q+1]
				}
				p--
				q++
			}
		}
	}
	return longest
}

// 查找最长的回文字符串
