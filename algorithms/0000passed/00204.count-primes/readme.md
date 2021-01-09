#### 题目
<p>统计所有小于非负整数&nbsp;<em>n&nbsp;</em>的质数的数量。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 10
<strong>输出:</strong> 4
<strong>解释:</strong> 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
</pre>


 #### 题解
 1、枚举
 ```go
func countPrimes(n int) (count int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return
}

// 判断是否是质数
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```
 时间复杂度O(n*sqrt(n)),空间复杂度O(1)
 
 2、埃氏筛
 如果x是质数，那么大于x的倍数2x,3x...一定不是质数。
 ```go
func countPrimes(n int) (count int) {
	isPrime := make([]bool,n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := 2*i; j < n; j+=i {
				isPrime[j] = false
			}
		}
	}
	return
}
```
 时间复杂度O(nlogn),空间复杂度O(n)
 
 3、线性筛
 埃氏筛中例如35会由5和7重复赋值，而线性筛中维护一个质数表，只标记质数中与x相乘的数
 ```go
func countPrimes(n int) int {
	isPrime := make([]bool,n)
	primes := []int{}
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
```
 时间复杂度O(n),空间复杂度O(n)
 