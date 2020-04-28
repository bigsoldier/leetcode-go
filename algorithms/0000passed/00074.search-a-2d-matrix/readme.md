#### 题目
<p>编写一个高效的算法来判断&nbsp;<em>m</em> x <em>n</em>&nbsp;矩阵中，是否存在一个目标值。该矩阵具有如下特性：</p>

<ul>
	<li>每行中的整数从左到右按升序排列。</li>
	<li>每行的第一个整数大于前一行的最后一个整数。</li>
</ul>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong>
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong>
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
<strong>输出:</strong> false</pre>


 #### 题解
 1、暴力法
 因为矩阵是增序的，那么将矩阵转为数组然后再用二分法查找
 ```go
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	var result []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = append(result, matrix[i][j])
		}
	}

	left,right := 0,len(result)
	for left < right {
		mid := (left + right)/2
		if target == result[mid] {
			return true
		}
		if target < result[mid] {
			right = mid
		}
		if target > result[mid] {
			left = mid
		}
	}
	return false
}
```
 结果超时
 
 2、改进
 其实数组没有必要，可以用索引找到他们
 ```go
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	left,right := 0,m*n-1
	for left <= right {
		mid := (left + right)/2
		if target == matrix[mid/n][mid%n] {
			return true
		}
		if target < matrix[mid/n][mid%n] {
			right = mid-1
		}
		if target > matrix[mid/n][mid%n] {
			left = mid+1
		}
	}
	return false
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-26/007401.png)
 时间复杂度O(log(mn)),空间复杂度O(1)