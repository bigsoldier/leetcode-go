#### 题目
<p>给定一个整数 <em>n</em>，求以&nbsp;1 ...&nbsp;<em>n</em>&nbsp;为节点组成的二叉搜索树有多少种？</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 3
<strong>输出:</strong> 5
<strong>解释:
</strong>给定 <em>n</em> = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3</pre>


 #### 题解
 1、[卡塔兰数](https://baike.baidu.com/item/%E5%8D%A1%E7%89%B9%E5%85%B0%E6%95%B0/6125746?fromtitle=%E5%8D%A1%E5%A1%94%E5%85%B0%E6%95%B0&fromid=9133402&fr=aladdin)
 C~0~=1, c~n+1~=2(2n+1)C~n~/n+2
 ```go
func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	return (4*n+2)*numTrees(n-1)/(n+2)
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-08/009601.png)
 时间复杂度O(n),空间复杂度O(1)
 
 2、动态规划
 遍历1...n,将其中一个数字为跟,1...(i-1)为左子树,(i+1)...n为右子树,由于根不同,每棵二叉树保证是单一的.
```go
func numTrees(n int) int {
	var dp = make([]int,n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
} 
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-08/009602.png)
 时间复杂度O(n^2^),空间复杂度O(n)