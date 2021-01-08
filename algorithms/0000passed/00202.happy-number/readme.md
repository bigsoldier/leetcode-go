#### 题目
<p>编写一个算法来判断一个数是不是&ldquo;快乐数&rdquo;。</p>

<p>一个&ldquo;快乐数&rdquo;定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。</p>

<p><strong>示例:&nbsp;</strong></p>

<pre><strong>输入:</strong> 19
<strong>输出:</strong> true
<strong>解释: 
</strong>1<sup>2</sup> + 9<sup>2</sup> = 82
8<sup>2</sup> + 2<sup>2</sup> = 68
6<sup>2</sup> + 8<sup>2</sup> = 100
1<sup>2</sup> + 0<sup>2</sup> + 0<sup>2</sup> = 1
</pre>


 #### 题解
 |  |   |   |
 | ---- | ---- | ---- | ---- |
 | 1 | 9 | 81 |
 | 2 | 99 | 162 |
 | 3 | 999 | 243 |
 | 4 | 9999 | 324 |
 | 13 | 9999999999999 | 1053
 
 我们发现对于3位数，它不可能大于243。对于4位及4位以上的数字，他们最终会降到3位数。
 
 ```go
func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; {
		n,m[n] = step(n),true
	}
	return n == 1
}

func step(n int) int {
	var sum int
	for n > 0 {
		sum += (n%10) * (n%10)
		n /= 10
	}
	return sum
}
```

 快慢指针
 ```go
func isHappy(n int) bool {
	slow,fast := n,step(n)
	for fast != 1 && slow != fast {
		slow,fast = step(slow),step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	var sum int
	for n > 0 {
		sum += (n%10) * (n%10)
		n /= 10
	}
	return sum
}
```