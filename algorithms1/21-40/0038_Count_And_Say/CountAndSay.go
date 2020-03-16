package _038_Count_And_Say

import "strconv"

//The count-and-say sequence is the sequence of integers with the first five ter
//ms as following:
//
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//
//
// 1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//
// Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-s
//ay sequence. You can do so recursively, in other words from the previous member
//read off the digits, counting the number of digits in groups of the same digit.
//
//
// Note: Each term of the sequence of integers will be represented as a string.
//
//
//
//
// Example 1:
//
//
//Input: 1
//Output: "1"
//Explanation: This is the base case.
//
//
// Example 2:
//
//
//Input: 4
//Output: "1211"
//Explanation: For n = 3 the term was "21" in which we have two groups "2" and "
//1", "2" can be read as "12" which means frequency = 1 and value = 2, the same wa
//y "1" is read as "11", so the answer is the concatenation of "12" and "11" which
// is "1211".
//
// Related Topics String

//leetcode submit region begin(Prohibit modification and deletion)
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ret := countAndSay(n - 1)

	var result string
	var count = 1
	for i := 1; i < len(ret); i++ {
		if ret[i] == ret[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + ret[i-1:i]
			count = 1
		}
	}
	result += strconv.Itoa(count) + ret[len(ret)-1:]
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
