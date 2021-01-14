#### 题目
<p>在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入: 
</strong>
1 0 1 0 0
1 0 <strong>1 1</strong> 1
1 1 <strong>1 1 </strong>1
1 0 0 1 0

<strong>输出: </strong>4</pre>


 #### 题解
 1、暴力法
 以左上角为顶点不断延伸，判断是否构成正方形
 ```go
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxSide int
	m,n := len(matrix),len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide,1)
				currSide := min(m-i,n-j)
				for k := 0; k < currSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for l := 0; l < k; l++ {
						if matrix[i+k][j+l]=='0' || matrix[i+l][j+k]=='0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide,k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide*maxSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(mn*min(m,n)^2^),空间复杂度O(1)
 
 2、动态规划
 一个4*4的正方形如果以右下角为顶点，去掉右下角可以重新形成3个3*3的正方形。
 
 以此为基础，dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
 ```go
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxSide int
	var dp = make([][]int,len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int,len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				maxSide = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide*maxSide
}

func min(a... int) (ans int) {
	ans = a[0]
	for _, i := range a {
		if i < ans {
			ans = i
		}
	}
	return ans
}
```
 时间复杂度O(mn),空间复杂度O(mn)