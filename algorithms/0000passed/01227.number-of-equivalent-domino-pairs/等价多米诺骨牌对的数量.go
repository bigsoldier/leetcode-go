package code

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	var dominoeSlice [100]int
	for _, dominoe := range dominoes {
		if dominoe[0] > dominoe[1] {
			dominoe[0], dominoe[1] = dominoe[1], dominoe[0]
		}
		num := dominoe[0]*10 + dominoe[1]
		ans += dominoeSlice[num]
		dominoeSlice[num]++
	}
	return ans
}
