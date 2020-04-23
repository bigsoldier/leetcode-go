package code

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	wordLen := len(words)
	words = append(words, "")
	var line []string
	var count int
	var result []string
	for i := 0; i < wordLen; i++ {
		line = append(line, words[i])
		count += len(words[i])
		if count+len(line)+len(words[i+1]) > maxWidth {
			idleSpaceCount := maxWidth - count
			if len(line) == 1 {
				result = append(result, line[0]+fillSpace(idleSpaceCount))
			} else {
				content := line[0]
				spaces := allocateSpace(idleSpaceCount, len(line)-1)
				for i := 1; i < len(line); i++ {
					content += fillSpace(spaces[i-1]) + line[i]
				}
				result = append(result, content)
			}

			// 置0
			line = nil
			count = 0
		} else if i == wordLen-1 {
			// 如果是最后一条，需要左对齐
			content := strings.Join(line, " ")
			content += fillSpace(maxWidth - len(content))
			result = append(result, content)
		}
	}
	return result
}

// 填充空格
func fillSpace(count int) string {
	var space string
	for i := 0; i < count; i++ {
		space += " "
	}
	return space
}

func allocateSpace(all, length int) []int {
	if length == 0 {
		return []int{all}
	}
	var slice = make([]int, length)
	a, b := all/length, all%length
	for i := 0; i < length; i++ {
		if b > 0 {
			slice[i] = a + 1
		} else {
			slice[i] = a
		}
		b--
	}
	return slice
}
