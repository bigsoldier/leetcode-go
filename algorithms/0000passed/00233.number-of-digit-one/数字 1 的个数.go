package code

func countDigitOne(n int) (ans int) {
	for i := 1; i <= n; i *= 10 {
		div := i * 10
		// (n/div)*i 低位是1的个数: 11,21,31
		// min(max(n%div-i+1,0),i)) 高位是1的个数(最多根据位数来): 10,11,12
		ans += (n/div)*i + min(max(n%div-i+1, 0), i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
