/**
 *Created by Xie Jian on 2020/2/14 13:00
 */
package _022_Generate_Parentheses

//
//Given n pairs of parentheses, write a function to generate all combinations of
// well-formed parentheses.
//
//
//
//For example, given n = 3, a solution set is:
//
//
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
// Related Topics String Backtracking

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var result = make([][]string, n+1)
	result[0] = []string{""}   // 0组括号的情况
	result[1] = []string{"()"} // 1组括号的情况
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ { // 遍历p q，其中p+q=i-1
			for _, k1 := range result[j] {
				for _, k2 := range result[i-1-j] {
					result[i] = append(result[i], "("+string(k1)+")"+string(k2))
				}
			}
		}
	}
	return result[n]
}

//leetcode submit region end(Prohibit modification and deletion)
