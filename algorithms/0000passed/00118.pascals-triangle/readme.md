#### 题目
<p>给定一个非负整数&nbsp;<em>numRows，</em>生成杨辉三角的前&nbsp;<em>numRows&nbsp;</em>行。</p>

<p><img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif"></p>

<p><small>在杨辉三角中，每个数是它左上方和右上方的数的和。</small></p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 5
<strong>输出:</strong>
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]</pre>


 #### 题解
 ```go
func generate(numRows int) [][]int {
	var results = make([][]int,numRows)
	for i := 0; i < numRows; i++ {
		results[i] = make([]int,i+1)
		results[i][0] = 1
		results[i][i] = 1
		for j := 1; j < i; j++ {
			results[i][j] = results[i-1][j-1]+results[i-1][j]
		}
	}
	return results
}
```
 时间复杂度O(n^2^2),空间复杂度O(n^2^),如果不考虑返回值的占用就是O(1)