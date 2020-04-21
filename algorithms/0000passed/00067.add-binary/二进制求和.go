package code

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	c := byte(0)
	var result string
	for i >= 0 || j >= 0 || c > 0 {
		sum := c
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		c = sum / 2
		result = string('0'+sum%2) + result
	}
	return result
}
