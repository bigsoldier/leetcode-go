package code

func isAnagram(s string, t string) bool {
	var char1, char2 [26]int
	for _, ch := range s {
		char1[ch-'a']++
	}
	for _, ch := range t {
		char2[ch-'a']++
	}
	return char1 == char2
}
