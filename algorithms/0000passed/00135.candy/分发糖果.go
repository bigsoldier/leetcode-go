package code

func candy(ratings []int) int {
	// prev是上一个同学分配到的，ans是总数
	prev, ans, inc, dec := 1, 1, 1, 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				prev = 1
			} else {
				prev++
			}
			ans += prev
			inc = prev
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			prev = 1
		}
	}
	return ans
}
