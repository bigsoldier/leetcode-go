package code

var (
	letterMap = []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	result []string
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	getString(digits, "")
	return result
}

func getString(digits, s string) {
	if len(digits) == 0 {
		result = append(result, s) // 最后取所有的字母
		return
	}

	letter := letterMap[digits[0]-'0']
	for i := 0; i < len(letter); i++ {
		getString(digits[1:], s+string(letter[i]))
	}
}
