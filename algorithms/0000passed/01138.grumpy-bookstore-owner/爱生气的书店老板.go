package code

func maxSatisfied(customers []int, grumpy []int, X int) int {
	var total int
	n := len(customers)
	for i := range customers {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := X; i < n; i++ {
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
