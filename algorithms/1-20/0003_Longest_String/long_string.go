package leetcode

//Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
//
// Related Topics Hash Table Two Pointers String Sliding Window
func lengthOfLongestSubstring(s string) int {
	var count = make(map[rune]int) // 存放字符的位置
	var max int
	left := 0
	for k, v := range s {
		if val, ok := count[v]; ok {
			if left < val {
				left = val
			}
		}
		if max < k-left+1 {
			max = k - left + 1
		}
		count[v] = k + 1 // 更新位置
	}
	return max
}

// 看了很多，确实类似于滑动窗口的方法更简单，有点类似于快速排序法
// 选取锚点，通过map统计字符出现的位置，如果出现过，则把锚点移动，在这个过程中更新最大长度
