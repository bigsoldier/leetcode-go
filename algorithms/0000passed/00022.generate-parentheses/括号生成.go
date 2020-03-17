package code

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
