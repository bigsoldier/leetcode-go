/**
 *Created by Xie Jian on 2020/2/18 16:39
 */
package _030_Substring_With_Concatenation_Of_All_Words

//You are given a string, s, and a list of words, words, that are all of the sam
//e length. Find all starting indices of substring(s) in s that is a concatenation
// of each word in words exactly once and without any intervening characters.
//
//
//
// Example 1:
//
//
//Input:
//  s = "barfoothefoobarman",
//  words = ["foo","bar"]
//Output: [0,9]
//Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" re
//spectively.
//The output order does not matter, returning [9,0] is fine too.
//
//
// Example 2:
//
//
//Input:
//  s = "wordgoodgoodgoodbestword",
//  words = ["word","good","best","word"]
//Output: []
//
// Related Topics Hash Table Two Pointers String

//leetcode submit region begin(Prohibit modification and deletion)
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var res []int
	wordLen := len(words[0])
	wordNum := len(words)
	if len(s) < wordLen {
		return nil
	}

	var mapWords = make(map[string]int)
	for _, word := range words {
		mapWords[word]++
	}

	for i := 0; i < wordLen; i++ { // 每wordLen长度间隔查询
		left, right, count := i, i, 0
		var tmpMap = make(map[string]int)
		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen
			if _, ok := mapWords[word]; !ok {
				count = 0
				left = right
				tmpMap = make(map[string]int)
			} else {
				tmpMap[word]++
				count++
				for tmpMap[word] > mapWords[word] {
					count--
					tmpMap[s[left:left+wordLen]]--
					left += wordLen
				}
				if count == wordNum {
					res = append(res, left)
				}
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
