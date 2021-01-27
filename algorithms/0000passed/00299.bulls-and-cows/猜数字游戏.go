package code

import "strconv"

func getHint(secret string, guess string) string {
	var count [10]int
	var A, B int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			A++
		} else {
			if count[secret[i]-'0'] < 0 {
				B++
			}
			count[secret[i]-'0']++
			if count[guess[i]-'0'] > 0 {
				B++
			}
			count[guess[i]-'0']--
		}
	}
	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
}
