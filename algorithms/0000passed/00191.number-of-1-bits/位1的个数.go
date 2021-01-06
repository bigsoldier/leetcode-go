package code

func hammingWeight(num uint32) int {
	var bits = 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			bits++
		}
		mask <<= 1
	}
	return bits
}
