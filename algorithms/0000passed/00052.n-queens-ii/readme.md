#### 题目
<p><em>n&nbsp;</em>皇后问题研究的是如何将 <em>n</em>&nbsp;个皇后放置在 <em>n</em>&times;<em>n</em> 的棋盘上，并且使皇后彼此之间不能相互攻击。</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/8-queens.png" style="height: 276px; width: 258px;"></p>

<p><small>上图为 8 皇后问题的一种解法。</small></p>

<p>给定一个整数 <em>n</em>，返回 <em>n</em> 皇后不同的解决方案的数量。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 4
<strong>输出:</strong> 2
<strong>解释:</strong> 4 皇后问题存在如下两个不同的解法。
[
&nbsp;[&quot;.Q..&quot;, &nbsp;// 解法 1
&nbsp; &quot;...Q&quot;,
&nbsp; &quot;Q...&quot;,
&nbsp; &quot;..Q.&quot;],

&nbsp;[&quot;..Q.&quot;, &nbsp;// 解法 2
&nbsp; &quot;Q...&quot;,
&nbsp; &quot;...Q&quot;,
&nbsp; &quot;.Q..&quot;]
]
</pre>


 #### 题解
 回溯法
 与51题一样，只是不需要将解集返回，只需要计数即可
 ```go
func totalNQueens(n int) int {
	col,diagonal1,diagonal2,row,result := make([]bool,n),make([]bool,2*n-1),make([]bool,2*n-1),[]int{},0
	putQueen(n,0,&col,&diagonal1,&diagonal2,&row,&result)
	return result
}

func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, result *int) {
	if index == n {
		*result++
		return
	}
	for i := 0; i < n; i++ {
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i]=true
			(*dia2)[index-i+n-1]=true
			putQueen(n,index+1,col,dia1,dia2,row,result)
			(*col)[i] = false
			(*dia1)[index+i]=false
			(*dia2)[index-i+n-1]=false
			*row=(*row)[:len(*row)-1]
		}
	}
	return
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-31/005201.png)