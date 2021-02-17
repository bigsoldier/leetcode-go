#### 题目
<p>在MATLAB中，有一个非常有用的函数 <code>reshape</code>，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据。</p>

<p>给出一个由二维数组表示的矩阵，以及两个正整数<code>r</code>和<code>c</code>，分别表示想要的重构的矩阵的行数和列数。</p>

<p>重构后的矩阵需要将原始矩阵的所有元素以相同的<strong>行遍历顺序</strong>填充。</p>

<p>如果具有给定参数的<code>reshape</code>操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> 
nums = 
[[1,2],
 [3,4]]
r = 1, c = 4
<strong>输出:</strong> 
[[1,2,3,4]]
<strong>解释:</strong>
行遍历nums的结果是 [1,2,3,4]。新的矩阵是 1 * 4 矩阵, 用之前的元素值一行一行填充新矩阵。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> 
nums = 
[[1,2],
 [3,4]]
r = 2, c = 4
<strong>输出:</strong> 
[[1,2],
 [3,4]]
<strong>解释:</strong>
没有办法将 2 * 2 矩阵转化为 2 * 4 矩阵。 所以输出原矩阵。
</pre>

<p><strong>注意：</strong></p>

<ol>
	<li>给定矩阵的宽和高范围在 [1, 100]。</li>
	<li>给定的 r 和 c 都是正数。</li>
</ol>


 #### 题解
 1、替代
 ```go
func matrixReshape(nums [][]int, r int, c int) [][]int {
    m,n := len(nums),len(nums[0])
    if m*n != r*c {
        return nums
    }
    var graph = make([][]int,r)
    var k int // 表示所在的行序
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            graph[k] = append(graph[k],nums[i][j])
            if len(graph[k]) == c {
                k++
            }
        }
    }
    return graph
}
```
 时间复杂度O(rc),空间复杂度O(1)
 
 2、映射
 我们把原二维数组“压扁”，变成一维数组
 $$
 (i,j) -> i*m+n
 $$
 然后再将一维数组变成二维数组。
 $$
 i = x/n
 j = x%n
 $$
 
 ```go
func matrixReshape(nums [][]int, r int, c int) [][]int {
    m,n := len(nums),len(nums[0])
    if m*n != r*c {
        return nums
    }
    var ans = make([][]int,r)
    for i := range ans {
        ans[i] = make([]int,c)
    }
    for i := 0;i<m*n;i++ {
        ans[i/c][i%c] = nums[i/n][i%n]
    }
    return ans
}
```
 时间复杂度O(rc),空间复杂度O(1)