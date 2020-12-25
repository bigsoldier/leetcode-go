package code

import "strings"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	var ret string
	for i := 0; i < len(s); i++ {
		if isAlnum(s[i]) {
			ret += string(s[i])
		}
	}
	ret = strings.ToLower(ret)
	for i := 0; i < len(ret)/2; i++ {
		if ret[i] != ret[len(ret)-i-1] {
			return false
		}
	}
	return true
}

func isAlnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
