package code

func largeGroupPositions(S string) (ans [][]int) {
	var index = 1
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			index++
		} else {
			if index > 2 {
				ans = append(ans, []int{i - index, i - 1})
			}
			index = 1
		}
	}
	if index > 2 {
		ans = append(ans, []int{len(S) - index, len(S) - 1})
	}
	return
}
