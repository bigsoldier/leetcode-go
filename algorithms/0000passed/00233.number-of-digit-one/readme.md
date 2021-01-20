#### 题目
<p>给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 13
<strong>输出:</strong> 6 
<strong>解释: </strong>数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。</pre>


 #### 题解
 1、暴力法
 遍历i从1到n
 ```go
func countDigitOne(n int) (ans int) {
	for i := 1; i <= n; i++ {
		ans += count(i)
	}
	return
}

func count(x int) (ans int) {
	for x > 0 {
		if x%10 == 1 {
			ans++
		}
		x /= 10
	}
	return
}
```
 时间复杂度O(nlogn)，空间复杂度O(logn)
 
 2、找规律
 我们可以观察到每10个数，个位上`1`会出现1次，每100个数，十位的`1`会出现一次。
 1,11,21,31,41,51,61...
 可以用 `n/(i*10)*i` 表示。
 
 如果十位是`1`,那么`1`总数x+1(x是个位上的数值). 如果十位上的数大于`1`，那么`1`的数量要加10。
 10,11,12,13,14,15,16,17...
 min(max((n mod(i*10))-i+1,0),i)
 
 例子: 有一个数 n=1234
 
 - 个位上`1`的个数 1234/10 + min(4,1) = 124
 - 十位上`1`的个数 1234/100*10 + min(21,10) = 130
 - 百位上`1`的个数 1234/1000*10 + min(135,100) = 200
 - 千位上`1`的个数 1234/10000*10 + min(235,1000) = 235
 ```go
func countDigitOne(n int) (ans int) {
	for i := 1; i <= n; i*=10 {
		div := i * 10
		// (n/div)*i 低位是1的个数: 11,21,31
		// min(max(n%div-i+1,0),i)) 高位是1的个数(最多根据位数来): 10,11,12
		ans += (n/div)*i + min(max(n%div-i+1,0),i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
``` 
 时间复杂度O(logn),空间复杂度O(1)