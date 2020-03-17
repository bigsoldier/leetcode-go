package code

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var temp int
	var y = x
	for y != 0 {
		temp = temp*10 + y%10
		y /= 10
	}
	if x == temp {
		return true
	}
	return false
}
