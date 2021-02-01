package code

func fairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	var mapA = make(map[int]struct{}, len(A))
	for _, a := range A {
		sumA += a
		mapA[a] = struct{}{}
	}
	for _, b := range B {
		sumB += b
	}
	diff := (sumA - sumB) / 2
	for _, b := range B {
		a := b + diff
		if _, ok := mapA[a]; ok {
			return []int{a, b}
		}
	}
	return nil
}
