package code

func countPrimes(n int) (count int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
