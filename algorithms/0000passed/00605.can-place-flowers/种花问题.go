package code

func canPlaceFlowers(flowerbed []int, n int) bool {
	count, index, length := 0, -1, len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			if index < 0 { // 只有[2,+]的时候才能种花
				count += i / 2
			} else {
				count += (i - index - 2) / 2
			}
			index = i
		}
	}
	// 计算最后一个区间
	if index < 0 {
		count += (length + 1) / 2
	} else {
		count += (length - index - 1) / 2
	}
	return count >= n
}
