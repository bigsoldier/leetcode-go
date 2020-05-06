package code

func combine(n int, k int) [][]int {
	var result [][]int
	var cur, used []int
	for i := 0; i < n; i++ {
		used = append(used, i+1)
	}
	add(&result, cur, used, k)
	return result
}

func add(result *[][]int, cur, used []int, k int) {
	if len(cur) == k {
		var c = make([]int, k)
		copy(c, cur)
		*result = append(*result, c)
		return
	}

	for i, u := range used {
		add(result, append(cur, u), used[i+1:], k)
	}
}
