package code

import "strings"

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	var ch2word = make(map[byte]string)
	var word2ch = make(map[string]byte)
	for i, word := range strs {
		ch := pattern[i]
		if (word2ch[word] > 0 && word2ch[word] != ch) || (ch2word[ch] != "" && ch2word[ch] != word) {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
