#### 题目
<p>编写一个高效的算法来搜索&nbsp;<em>m</em>&nbsp;x&nbsp;<em>n</em>&nbsp;矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：</p>

<ul>
	<li>每行的元素从左到右升序排列。</li>
	<li>每列的元素从上到下升序排列。</li>
</ul>

<p><strong>示例:</strong></p>

<p>现有矩阵 matrix 如下：</p>

<pre>[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
</pre>

<p>给定 target&nbsp;=&nbsp;<code>5</code>，返回&nbsp;<code>true</code>。</p>

<p>给定&nbsp;target&nbsp;=&nbsp;<code>20</code>，返回&nbsp;<code>false</code>。</p>


 #### 题解
 1、暴力法
 ```go
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
 ```
 时间复杂度O(mn),空间复杂度O(1)
 
 2、二分搜索
 对角线迭代搜索
 ```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	shorter := min(len(matrix),len(matrix[0]))
	for i := 0; i < shorter; i++ {
		if binarySearch(matrix, target, i, true) ||
			binarySearch(matrix, target, i, false) {
			return true
		}
	}
	return false
}

func binarySearch(matrix [][]int, target int, start int, vertical bool) bool {
	lo := start
	var hi int
	if vertical {
		hi = len(matrix[0])-1
	} else {
		hi = len(matrix)-1
	}

	for lo <= hi {
		mid := (lo+hi)/2
		if vertical { // 行查询
			if matrix[start][mid] < target {
				lo = mid + 1
			} else if matrix[start][mid] > target {
				hi = mid-1
			} else {
				return true
			}
		} else { // 列查询
			if matrix[mid][start] < target {
				lo = mid+1
			} else if matrix[mid][start] > target {
				hi = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 是啊金复杂度O(logn!),空间复杂度O(1)
 
 3、
 从矩阵中找到一个数，如果target比他小，那么他的左下区和右上区都是他的目标
 ```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return binarySearch(matrix,target,0,0,len(matrix)-1,len(matrix[0])-1)
}

func binarySearch(matrix [][]int, target int, x1,y1,x2,y2 int) bool {
	if x1 > x2 || y1 > y2 {
		return false
	} else if target < matrix[x1][y1] || target > matrix[x2][y2] {
		return false
	}
	mid := y1 + (y2-y1)/2
	// 找到matrix[row-1][mid]<matrix[row][mid]
	row := x1
	for row <= x2 && matrix[row][mid] <= target {
		if matrix[row][mid]==target {
			return true
		}
		row++
	}
	return binarySearch(matrix,target,row,y1,x2,mid-1) || binarySearch(matrix,target,x1,mid+1,row-1,y2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(nlogn),空间复杂度O(1)
 
 4、
 因为矩阵是有序的，所以我们应该根据某种顺序来查找。
 - 选左上角，往右走和往下走都是增大的，不知道要增大哪一个
 - 选右下角，往左走和往上走都是减小的，不知道要减小哪一个
 - 选左下角，往右走是增大，往上走是减小，可以根据数字来判断
 - 选右上角，往右走是减小，往上走是增大，可以根据数字来判断
 ```go
func searchMatrix(matrix [][]int, target int) bool {
	var i, j = len(matrix)-1,0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
```
 时间复杂度O(m+n),空间复杂度O(1)
 