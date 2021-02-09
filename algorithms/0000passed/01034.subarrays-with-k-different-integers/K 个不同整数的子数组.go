package code

func subarraysWithKDistinct(A []int, K int) (ans int) {
	n := len(A)
	num1, num2 := make(map[int]int, n+1), make(map[int]int, n+1)
	var left1, left2, to1, to2 int
	for _, a := range A {
		if num1[a] == 0 {
			to1++
		}
		num1[a]++
		if num2[a] == 0 {
			to2++
		}
		num2[a]++
		for to1 > K {
			num1[A[left1]]--
			if num1[A[left1]] == 0 {
				to1--
			}
			left1++
		}
		for to2 > K-1 {
			num2[A[left2]]--
			if num2[A[left2]] == 0 {
				to2--
			}
			left2++
		}
		ans += left2 - left1
	}
	return
}
