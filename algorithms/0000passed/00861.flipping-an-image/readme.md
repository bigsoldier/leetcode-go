#### 题目
<p>给定一个二进制矩阵&nbsp;<code>A</code>，我们想先水平翻转图像，然后反转图像并返回结果。</p>

<p>水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转&nbsp;<code>[1, 1, 0]</code>&nbsp;的结果是&nbsp;<code>[0, 1, 1]</code>。</p>

<p>反转图片的意思是图片中的&nbsp;<code>0</code>&nbsp;全部被&nbsp;<code>1</code>&nbsp;替换，&nbsp;<code>1</code>&nbsp;全部被&nbsp;<code>0</code>&nbsp;替换。例如，反转&nbsp;<code>[0, 1, 1]</code>&nbsp;的结果是&nbsp;<code>[1, 0, 0]</code>。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入: </strong>[[1,1,0],[1,0,1],[0,0,0]]
<strong>输出: </strong>[[1,0,0],[0,1,0],[1,1,1]]
<strong>解释:</strong> 首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
     然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入: </strong>[[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
<strong>输出: </strong>[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
<strong>解释:</strong> 首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
     然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
</pre>

<p><strong>说明:</strong></p>

<ul>
	<li><code>1 &lt;= A.length = A[0].length &lt;= 20</code></li>
	<li><code>0 &lt;= A[i][j]&nbsp;&lt;=&nbsp;1</code></li>
</ul>


 #### 题解
 最直观的做法就是将矩阵A的每一行水平翻转，然后对矩阵的每个元素再反转。
 这个过程需要对矩阵遍历两次。
 
 假设矩阵的行数和列数都是n，考虑下标left和right，其中left<right且left+right=n-1.
 当0<=i<n时，对i行进行水平翻转后，A[i][left]和A[i][right]的元素会互换，进行反转后，元素值都会改变。
 - A[i][left]=A[i][right]=0,水平翻转后,再反转,A[i][left]=A[i][right]=1
 - A[i][left]=A[i][right]=1,水平翻转后,再反转,A[i][left]=A[i][right]=0
 - A[i][left]=0,A[i][right]=1,水平翻转后,再反转,A[i][left]=0,A[i][right]=1
 - A[i][left]=1,A[i][right]=0,水平翻转后,再反转,A[i][left]=1,A[i][right]=0
 
 从这里可以看出，A[i][left]等于A[i][right]时，需要对两个值进行反转，如果不相等不进行任何操作。
 ```go
func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		left,right := 0,len(a)-1
		for left < right {
			if a[left] == a[right] {
				a[left] ^= 1
				a[right] ^= 1
			}
			left++;right--
		}
		if left == right {
			a[left] ^= 1
		}
	}
	return A
}
```
 时间复杂度O(n^2^),空间复杂度O(1)