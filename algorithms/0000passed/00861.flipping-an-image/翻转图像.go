package code

func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		left, right := 0, len(a)-1
		for left < right {
			if a[left] == a[right] {
				a[left] ^= 1
				a[right] ^= 1
			}
			left++
			right--
		}
		if left == right {
			a[left] ^= 1
		}
	}
	return A
}
