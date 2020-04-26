#### 题目
<p>给定一个&nbsp;<em>m</em> x <em>n</em> 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用<strong><a href="http://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank">原地</a></strong>算法<strong>。</strong></p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 
[
&nbsp; [1,1,1],
&nbsp; [1,0,1],
&nbsp; [1,1,1]
]
<strong>输出:</strong> 
[
&nbsp; [1,0,1],
&nbsp; [0,0,0],
&nbsp; [1,0,1]
]
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 
[
&nbsp; [0,1,2,0],
&nbsp; [3,4,5,2],
&nbsp; [1,3,1,5]
]
<strong>输出:</strong> 
[
&nbsp; [0,0,0,0],
&nbsp; [0,4,5,0],
&nbsp; [0,3,1,0]
]</pre>

<p><strong>进阶:</strong></p>

<ul>
	<li>一个直接的解决方案是使用 &nbsp;O(<em>m</em><em>n</em>)&nbsp;的额外空间，但这并不是一个好的解决方案。</li>
	<li>一个简单的改进方案是使用 O(<em>m</em>&nbsp;+&nbsp;<em>n</em>) 的额外空间，但这仍然不是最好的解决方案。</li>
	<li>你能想出一个常数空间的解决方案吗？</li>
</ul>


 #### 题解
 1、额外空间
 记录出现的行列，然后再将出现的行列置0
 ```go
func setZeroes(matrix [][]int) {
	m,n := len(matrix),len(matrix[0])
	var rows,colume = make(map[int]int,m),make(map[int]int,n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = i
				colume[j] = j
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_,ok1 := rows[i]
			_,ok2 := colume[j]
			if ok1 || ok2 {
				matrix[i][j] = 0
			}
		}
	}
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-26/007301.png)
 时间复杂度O(mn),空间复杂度O(m+n)
 
 2、O(1)空间法
 将出现0的行列数据改为一个不正常的数，之后再遍历查到的更改为0
 ```go
func setZeroes(matrix [][]int) {
	m,n := len(matrix),len(matrix[0])
	var defineInf = math.MinInt16
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = defineInf
					}
				}
				for k := 0; k < n; k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = defineInf
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == defineInf {
				matrix[i][j] = 0
			}
		}
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-26/007302.png)
 时间复杂度O(mn(m+n)),空间复杂度O(1)
 
 3、改进的O(1)空间法
 方法2里重复对已经置0的数据重复置0,这个方法就是解决这个问题的.
 将出现的0的该行该列的第一个置0,然后根据首列和首行为0的去置0.
 首行和首列需要单独讨论
 ```go
func setZeroes(matrix [][]int) {
	m,n := len(matrix),len(matrix[0])
	var head bool

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 { // 如果首列有数据，则首列置0
			head = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j]==0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if head {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-26/007303.png)
 时间复杂度O(mn),空间复杂度O(1)