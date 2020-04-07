#### 题目
<p>给定一个 <em>n&nbsp;</em>&times;&nbsp;<em>n</em> 的二维矩阵表示一个图像。</p>

<p>将图像顺时针旋转 90 度。</p>

<p><strong>说明：</strong></p>

<p>你必须在<strong><a href="https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank">原地</a></strong>旋转图像，这意味着你需要直接修改输入的二维矩阵。<strong>请不要</strong>使用另一个矩阵来旋转图像。</p>

<p><strong>示例 1:</strong></p>

<pre>给定 <strong>matrix</strong> = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

<strong>原地</strong>旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
</pre>

<p><strong>示例 2:</strong></p>

<pre>给定 <strong>matrix</strong> =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

<strong>原地</strong>旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
</pre>


 #### 题解
 ```go
func rotate(matrix [][]int) {
	// 沿对角线转置
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}

	// 按行对称
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j],matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1],matrix[i][j]
		}
	}
}
 ```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-25/004801.png)
时间复杂度O(n^2^),空间复杂度O(1)(原地翻转)

2、矩阵旋转
除中心元素外，其实是四个元素的调换位置(可以找一个正方形画图来确定四个点的位置)
假设一个元素点的位置为(i,j),对应的其他三个点的位置为(n-1-j,i),(n-1-i,n-1-j),(j,n-1-i)
```go
func rotate(matrix [][]int) {
	// 四个点元素的切换
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			var tmp = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-25/004802.png)
时间复杂度O(n^2^),空间复杂度O(1)