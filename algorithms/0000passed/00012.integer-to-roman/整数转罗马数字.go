package code

func intToRoman(num int) string {
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var i int
	var Roman string

	for i < len(nums) {
		for num >= nums[i] {
			Roman += romans[i]
			num -= nums[i]
		}
		i++
	}
	return Roman
}
