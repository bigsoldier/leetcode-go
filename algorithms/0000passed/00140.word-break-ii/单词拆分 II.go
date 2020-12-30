package code

import "strings"

func wordBreak(s string, wordDict []string) (result []string) {
	wordSet := map[string]struct{}{}
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}
	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWord := range backtrack(i) {
					wordList = append(wordList, append([]string{word}, nextWord...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordList = append(wordList, []string{word})
		}
		dp[index] = wordList
		return wordList
	}
	for _, words := range backtrack(0) {
		result = append(result, strings.Join(words, " "))
	}
	return
}
