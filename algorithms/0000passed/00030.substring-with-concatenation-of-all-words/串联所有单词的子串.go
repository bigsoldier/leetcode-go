package code

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var res []int
	wordLen := len(words[0])
	wordNum := len(words)
	if len(s) < wordLen {
		return nil
	}

	var mapWords = make(map[string]int)
	for _, word := range words {
		mapWords[word]++
	}

	for i := 0; i < wordLen; i++ { // 每wordLen长度间隔查询
		left, right, count := i, i, 0
		var tmpMap = make(map[string]int)
		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen
			if _, ok := mapWords[word]; !ok {
				count = 0
				left = right
				tmpMap = make(map[string]int)
			} else {
				tmpMap[word]++
				count++
				for tmpMap[word] > mapWords[word] {
					count--
					tmpMap[s[left:left+wordLen]]--
					left += wordLen
				}
				if count == wordNum {
					res = append(res, left)
				}
			}
		}
	}

	return res
}
