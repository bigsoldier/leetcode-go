#### 题目
<p>给定一个非负索引&nbsp;<em>k</em>，其中 <em>k</em>&nbsp;&le;&nbsp;33，返回杨辉三角的第 <em>k </em>行。</p>

<p><img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif"></p>

<p><small>在杨辉三角中，每个数是它左上方和右上方的数的和。</small></p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 3
<strong>输出:</strong> [1,3,3,1]
</pre>

<p><strong>进阶：</strong></p>

<p>你可以优化你的算法到 <em>O</em>(<em>k</em>) 空间复杂度吗？</p>


 #### 题解
 1、递推关系
 ```go
func getRow(rowIndex int) []int {
    var rows = make([][]int,rowIndex+1)
    for i := 0; i < rowIndex+1;i++ {
        rows[i] = make([]int,i+1)
        rows[i][0],rows[i][i] = 1,1
        for j := 1;j < i;j++ {
            rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
        }
    }
    return rows[rowIndex]
}
 ```
 时间复杂度O(k^2^),空间复杂度O(k^2^)
 **优化**
 我们会发现计算i行时只会用到i-1行的数据，那么可以用到滚动数组的思想优化空间复杂度。
 ```go
func getRow(rowIndex int) []int {
    var pre,cur []int
    for i := 0; i <= rowIndex; i++ {
        cur = make([]int,i+1)
        cur[0],cur[i] = 1,1
        for j := 1; j < i;j++ {
            cur[j] = pre[j-1]+pre[j]
        }
        pre = cur
    }
    return pre
}
 ```
 时间复杂度O(k^2^),空间复杂度O(k)
 **再优化**
 我们能够发现计算j项元素和上一行的j项元素和j-1项元素相加，因此我们可以倒着计算当行。
 ```go
func getRow(rowIndex int) []int {
    var row = make([]int,rowIndex+1)
    row[0] = 1
    for i := 1; i <= rowIndex;i++ {
        for j := i;j > 0;j-- {
            row[j] += row[j-1]
        }
    }
    return row
}
 ```
 时间复杂度O(k^2^),空间复杂度O(k)

 2、数学公式
 可以得到同一行的相邻组合数的关系
$$
C^m_n = C^{m-1}_n \times \frac{n-m+1}{m}
$$
 ，由于`C_0_ = 1`.
 ```go
func getRow(rowIndex int) []int {
    var row = make([]int,rowIndex+1)
    row[0] = 1
    for i := 1;i <= rowIndex; i++ {
        row[i] = row[i-1]*(rowIndex-i+1)/i
    }
    return row
}
```
 时间复杂度O(k),空间复杂度O(k)