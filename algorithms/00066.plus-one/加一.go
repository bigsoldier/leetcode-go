package code

func plusOne(digits []int) []int {
	var multi = 1
	for i := 0; i < len(digits); i++ {
		multi *= 10
	}
	var ret int
	for i := 0; i < len(digits); i++ {
		multi /= 10
		ret += digits[i] * multi
	}
	ret += 1
	var result []int
	for ret > 0 {
		result = append(result, ret%10)
		ret /= 10
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
