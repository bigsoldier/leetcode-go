/**
 *Created by Xie Jian on 2020/2/12 15:29
 */
package _017_Letter_Combinations_Of_A_Phone_Number

//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
//
//
// Example:
//
//
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// Note:
//
// Although the above answer is in lexicographical order, your answer could be in any order you want.
// Related Topics String Backtracking

//leetcode submit region begin(Prohibit modification and deletion)
var (
	letterMap = []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	result []string
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	getString(digits, "")
	return result
}

func getString(digits, s string) {
	if len(digits) == 0 {
		result = append(result, s) // 最后取所有的字母
		return
	}

	letter := letterMap[digits[0]-'0']
	for i := 0; i < len(letter); i++ {
		getString(digits[1:], s+string(letter[i]))
	}
}

//leetcode submit region end(Prohibit modification and deletion)
