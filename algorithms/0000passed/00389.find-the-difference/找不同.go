package code

func findTheDifference(s string, t string) byte {
	var slice [26]int
	for _, val := range s {
		slice[val-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		slice[ch-'a']--
		if slice[ch-'a'] < 0 {
			return ch
		}
	}
}
