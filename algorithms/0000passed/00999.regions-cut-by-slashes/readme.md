#### 题目
<p>在由 1 x 1 方格组成的 N x N 网格&nbsp;<code>grid</code> 中，每个 1 x 1&nbsp;方块由 <code>/</code>、<code>\</code> 或空格构成。这些字符会将方块划分为一些共边的区域。</p>

<p>（请注意，反斜杠字符是转义的，因此 <code>\</code> 用 <code>&quot;\\&quot;</code>&nbsp;表示。）。</p>

<p>返回区域的数目。</p>

<p>&nbsp;</p>

<ol>
</ol>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：
</strong>[
&nbsp; &quot; /&quot;,
&nbsp; &quot;/ &quot;
]
<strong>输出：</strong>2
<strong>解释：</strong>2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/1.png"></pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：
</strong>[
&nbsp; &quot; /&quot;,
&nbsp; &quot;  &quot;
]
<strong>输出：</strong>1
<strong>解释：</strong>2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/2.png"></pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：
</strong>[
&nbsp; &quot;\\/&quot;,
&nbsp; &quot;/\\&quot;
]
<strong>输出：</strong>4
<strong>解释：</strong>（回想一下，因为 \ 字符是转义的，所以 &quot;\\/&quot; 表示 \/，而 &quot;/\\&quot; 表示 /\。）
2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/3.png"></pre>

<p><strong>示例 4：</strong></p>

<pre><strong>输入：
</strong>[
&nbsp; &quot;/\\&quot;,
&nbsp; &quot;\\/&quot;
]
<strong>输出：</strong>5
<strong>解释：</strong>（回想一下，因为 \ 字符是转义的，所以 &quot;/\\&quot; 表示 /\，而 &quot;\\/&quot; 表示 \/。）
2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/4.png"></pre>

<p><strong>示例 5：</strong></p>

<pre><strong>输入：
</strong>[
&nbsp; &quot;//&quot;,
&nbsp; &quot;/ &quot;
]
<strong>输出：</strong>3
<strong>解释：</strong>2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/5.png">
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= grid.length == grid[0].length &lt;= 30</code></li>
	<li><code>grid[i][j]</code> 是&nbsp;<code>&#39;/&#39;</code>、<code>&#39;\&#39;</code>、或&nbsp;<code>&#39; &#39;</code>。</li>
</ol>


 #### 题解
 1、深度优先
 /      \
 001    100
 010    010
 100    001
 ```go
func regionsBySlashes(grid []string) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// grid转化为矩阵
	m,n := len(grid),len(grid[0])
	var matrix = make([][]int,m*3)
	for i := 0; i < m*3; i++ {
		matrix[i] = make([]int,n*3)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '/' {
				matrix[i*3+2][j*3] = 1
				matrix[i*3+1][j*3+1] = 1
				matrix[i*3][j*3+2] = 1
			} else if grid[i][j] == '\\' {
				matrix[i*3][j*3] = 1
				matrix[i*3+1][j*3+1] = 1
				matrix[i*3+2][j*3+2] = 1
			}
		}
	}
	for i := 0; i < m*3; i++ {
		for j := 0; j < n*3; j++ {
			if matrix[i][j] == 0 {
				dfs(matrix,i,j)
				ans++
			}
		}
	}
	return
}

func dfs(matrix [][]int,i,j int) {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) || matrix[i][j] == 1 {
		return
	}
	matrix[i][j] = 1
	dfs(matrix,i-1,j)
	dfs(matrix,i+1,j)
	dfs(matrix,i,j-1)
	dfs(matrix,i,j+1)
}
 ```
 时间复杂度O(n^2^logn),空间复杂度O(n^2^)
 
 2、并查集
 我们将一个方块分为4个小三角形
 \/
 /\
 上右下左编号为0,1,2,3，
 
 单元格内的合并
 - 空：四个合并
 - /： 0和3合并，1和2合并
 - \:  0和1合并，2和3合并
 
 单元格间合并
 - 向左，向上
 - 向左，向下
 - 向右，向上
 - 向右，向下
 
 ```go
func regionsBySlashes(grid []string) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// grid转化为矩阵
	m := len(grid)

	cnt := 4*m*m
	fa := make([]int,cnt)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] =find(fa[x])
		}
		return fa[x]
	}
	merge := func(x,y int) {
		fx,fy := find(x),find(y)
		if fx != fy {
			fa[fx] = fy
			cnt--
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			idx := i*m+j
			// 单元格间的合并
			if i < m-1 { // 向下
				buttom := idx+m
				merge(idx*4+2,buttom*4) // 2和下的0合并
			}
			if j < m-1 { // 向右
				right := idx+1
				merge(idx*4+1,right*4+3) // 1和右的3合并
			}
			// 单元格内的合并
			if grid[i][j] == '/' {
				merge(idx*4,idx*4+3)
				merge(idx*4+1,idx*4+2)
			} else if grid[i][j] == '\\' {
				merge(idx*4,idx*4+1)
				merge(idx*4+2,idx*4+3)
			} else {
				merge(idx*4,idx*4+1)
				merge(idx*4,idx*4+2)
				merge(idx*4,idx*4+3)
			}
		}
	}
	return cnt
}
```