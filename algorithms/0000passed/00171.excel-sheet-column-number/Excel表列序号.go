package code

func titleToNumber(s string) (ans int) {
	for i := 0; i < len(s); i++ {
		ans = ans*26 + int(s[i]-'A') + 1
	}
	return
}
