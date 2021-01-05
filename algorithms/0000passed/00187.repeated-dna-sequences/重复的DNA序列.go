package code

func findRepeatedDnaSequences(s string) (ans []string) {
	const L = 10
	m := map[string]int{}
	for i := 0; i < len(s)-L+1; i++ {
		if v := m[s[i:i+L]]; v == 1 {
			ans = append(ans, s[i:i+L])
		}
		m[s[i:i+L]]++
	}
	return
}
