package code

func prefixesDivBy5(A []int) []bool {
	var ans = make([]bool, len(A))
	x := 0
	for i, v := range A {
		x = (x<<1 | v) % 5
		ans[i] = x == 0
	}
	return ans
}
