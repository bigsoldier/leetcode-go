package code

func singleNumber(nums []int) []int {
	var xor int
	// 求出 x^y
	for _, num := range nums {
		xor ^= num
	}
	// 求出最右边的1, 这里分不清哪个是x，哪个是y
	lowest := xor & -xor
	var a int
	for _, num := range nums {
		//
		if num&lowest == 0 {
			a ^= num // 通过异或来找出a
		}
	}
	return []int{a, xor ^ a}
}
