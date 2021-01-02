package code

func reverseWords(s string) (ret string) {
	ans := []string{""}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans = append(ans, "")
		} else {
			ans[len(ans)-1] += string(s[i])
		}
	}
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] != "" {
			ret += ans[i] + " "
		}
	}
	return ret[:len(ret)-1]
}
