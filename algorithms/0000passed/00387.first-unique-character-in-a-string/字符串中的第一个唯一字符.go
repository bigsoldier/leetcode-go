package code

func firstUniqChar(s string) int {
	var slice = make([]int, 26)
	for i := 0; i < len(s); i++ {
		slice[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if slice[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
