package code

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ret := countAndSay(n - 1)

	var result string
	var count = 1
	for i := 1; i < len(ret); i++ {
		if ret[i] == ret[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + ret[i-1:i]
			count = 1
		}
	}
	result += strconv.Itoa(count) + ret[len(ret)-1:]
	return result
}
