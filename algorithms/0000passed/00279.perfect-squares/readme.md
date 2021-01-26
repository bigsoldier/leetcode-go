#### 题目
<p>给定正整数&nbsp;<em>n</em>，找到若干个完全平方数（比如&nbsp;<code>1, 4, 9, 16, ...</code>）使得它们的和等于<em> n</em>。你需要让组成和的完全平方数的个数最少。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> <em>n</em> = <code>12</code>
<strong>输出:</strong> 3 
<strong>解释: </strong><code>12 = 4 + 4 + 4.</code></pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> <em>n</em> = <code>13</code>
<strong>输出:</strong> 2
<strong>解释: </strong><code>13 = 4 + 9.</code></pre>


 #### 题解
 1、暴力法
 ```go
func numSquares(n int) (cnt int) {
	var square []int
	for i := 1; i*i <= n; i++ {
		square = append(square, i*i)
	}
	return minNumSquares(square,n)
}

func minNumSquares(square []int, num int) (ans int) {
	// 找到平方数的情况
	for _, v := range square {
		if num == v {
			return 1
		}
	}
	ans = math.MaxInt32
	// 找到可能解中最小的值
	for _, v := range square {
		if v > num {
			break
		}
		newNum := minNumSquares(square,num-v)+1
		ans = min(ans,newNum)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 
 2、动态规划
 暴力超时的原因很简单，因为我们重复计算了中间解。
 
 numSquares[n] = min(numSquares(n-k)+1)
 ```go
func numSquares(n int) (cnt int) {
	var square []int
	for i := 1; i*i <= n; i++ {
		square = append(square, i*i)
	}
	var dp = make([]int,n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		for _, v := range square {
			if i < v {
				break
			}
			dp[i] = min(dp[i],dp[i-v]+1)
		}
	}
	return dp[n]
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(n*n^1/2^),空间复杂度O(n)
 
 3、贪心枚举
 ```go
func numSquares(n int) (cnt int) {
	var square []int
	for i := 1; i*i <= n; i++ {
		square = append(square, i*i)
	}
	for cnt = 1; cnt <= n; cnt++ {
		if isDivided(square, n, cnt) {
			return cnt
		}
	}
	return
}

func isDivided(square []int, n int, count int) bool {
	if count == 1 {
		for _, v := range square {
			if v == n {
				return true
			}
		}
		return false
	}

	for _, v := range square {
		if isDivided(square, n-v, count-1) {
			return true
		}
	}
	return false
}
```
 时间复杂度O(n^h/2^),空间复杂度O(n^1/2^),h是树的高度
 
 4、贪心+BFS
 ```go
func numSquares(n int) (cnt int) {
	var square []int
	for i := 1; i*i <= n; i++ {
		square = append(square, i*i)
	}
	queue := []int{n}
	for len(queue) > 0 {
		cnt++
		var slice []int
		for _, q := range queue {
			for _, num := range square {
				if q == num {
					return cnt
				} else if q < num {
					break
				} else {
					slice = append(slice, q-num)
				}
			}
		}
		queue = slice
	}
	return
}
```
 时间复杂度O(n^h/2^),空间复杂度O(n^h/2^),h是树的高度
 
 5、数学运算
 [四平方数和定理](https://zh.wikipedia.org/wiki/%E5%9B%9B%E5%B9%B3%E6%96%B9%E5%92%8C%E5%AE%9A%E7%90%86)
 
 每个自然数可以表示成四个整数的平方和。
 
 - 检查数字是否为 n = 4^k^(8m+7)，如果是返回4
 - 检查是否是或可以拆分为完全平方数
 - 不能则返回3
 ```go
func numSquares(n int) (cnt int) {
	for n &3 == 0 {
		n >>= 2 // 4^k
	}
	if n&7 == 7 { // mod 8 
		return 4
	}
	if isSquare(n) {
		return 1
	}
	for i := 1; i <= int(math.Pow(float64(n), 0.5)+1); i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isSquare(n int) bool {
	sq := math.Sqrt(float64(n))
	return int(sq) * int(sq) == n
}
```