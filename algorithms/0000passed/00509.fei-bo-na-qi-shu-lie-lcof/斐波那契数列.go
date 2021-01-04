package code

func fib(n int) int {
	f0, f1 := 0, 1
	if n == 0 {
		return f0
	}
	if n == 1 {
		return f1
	}
	for i := 2; i <= n; i++ {
		tmp := f1
		f1 = f0 + f1
		f0 = tmp
	}
	return f1
}
