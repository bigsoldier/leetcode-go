package code

func canCompleteCircuit(gas []int, cost []int) int {
	var remains, debt, start int
	for i := 0; i < len(gas); i++ {
		remains += gas[i] - cost[i]
		if remains < 0 {
			start = i + 1
			debt += remains
			remains = 0
		}
	}
	if remains+debt < 0 {
		return -1
	}
	return start
}
