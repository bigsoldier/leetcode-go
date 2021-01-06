package code

func reverseBits(num uint32) uint32 {
	var ret = uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power--
	}
	return ret
}
