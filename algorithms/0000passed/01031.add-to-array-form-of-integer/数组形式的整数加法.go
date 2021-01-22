package code

func addToArrayForm(A []int, K int) (ans []int) {
	for i := len(A) - 1; i >= 0; i-- {
		sum := A[i] + K%10
		K /= 10
		if sum >= 10 {
			K++
			sum -= 10
		}
		ans = append(ans, sum)

	}
	for ; K > 0; K /= 10 {
		ans = append(ans, K%10)
	}
	reverse(ans)
	return
}

func reverse(A []int) {
	n := len(A)
	for i := 0; i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}
}
