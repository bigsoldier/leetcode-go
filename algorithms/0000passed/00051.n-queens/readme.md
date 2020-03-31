#### 题目
<p><em>n&nbsp;</em>皇后问题研究的是如何将 <em>n</em>&nbsp;个皇后放置在 <em>n</em>&times;<em>n</em> 的棋盘上，并且使皇后彼此之间不能相互攻击。</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/8-queens.png"></p>

<p><small>上图为 8 皇后问题的一种解法。</small></p>

<p>给定一个整数 <em>n</em>，返回所有不同的&nbsp;<em>n&nbsp;</em>皇后问题的解决方案。</p>

<p>每一种解法包含一个明确的&nbsp;<em>n</em> 皇后问题的棋子放置方案，该方案中 <code>&#39;Q&#39;</code> 和 <code>&#39;.&#39;</code> 分别代表了皇后和空位。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 4
<strong>输出:</strong> [
 [&quot;.Q..&quot;,  // 解法 1
  &quot;...Q&quot;,
  &quot;Q...&quot;,
  &quot;..Q.&quot;],

 [&quot;..Q.&quot;,  // 解法 2
  &quot;Q...&quot;,
  &quot;...Q&quot;,
  &quot;.Q..&quot;]
]
<strong>解释:</strong> 4 皇后问题存在两个不同的解法。
</pre>


 #### 题解
 1、暴力法
 列出棋盘上摆放皇后的所有可能性，然后再比较是否是皇后问题的解决方案。时间复杂度O(n^n^)
 
 2、回溯法
 **一行只能有一个皇后且一列也只能有一个皇后**
 **对于所有次对角线有 行号+列号=常数(/), 所有主对角线有 行号-列号=常数(\)**
 ```go
func solveNQueens(n int) [][]string {
	col,diagonal1,diagonal2,row,result := make([]bool,n),make([]bool,2*n-1),make([]bool,2*n-1),[]int{},[][]string{}
	putQueen(n,0,&col,&diagonal1,&diagonal2,&row,&result)
	return result
}

func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, result *[][]string) {
	if index == n {
		*result = append(*result, generateBoard(n,row))
		return
	}
	for i := 0; i < n; i++ {
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i]=true
			(*dia2)[index-i+n-1]=true
			putQueen(n,index+1,col,dia1,dia2,row,result)
			// 再将数据还原
			(*col)[i] = false
			(*dia1)[index+i]=false
			(*dia2)[index-i+n-1]=false
			*row=(*row)[:len(*row)-1]
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	var board  []string
	var result string
	for i := 0; i < n; i++ {
		result += "."
	}
	for i := 0; i<n; i++ {
		board = append(board, result)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-31/005101.png)