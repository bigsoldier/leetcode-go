#### 题目
<p>给定两个整数 <em>n</em> 和 <em>k</em>，返回 1 ... <em>n </em>中所有可能的 <em>k</em> 个数的组合。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;n = 4, k = 2
<strong>输出:</strong>
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]</pre>


 #### 题解
 很熟悉的一套dfs深度搜索加剪枝
 ```go
func combine(n int, k int) [][]int {
	var result [][]int
	var cur,used []int
	for i := 0; i < n; i++ {
		used = append(used, i+1)
	}
	add(&result,cur,used,k)
	return result
}

func add(result *[][]int, cur, used []int, k int) {
	if len(cur) == k {
		var c = make([]int,k)
		copy(c, cur)
		*result = append(*result, c)
		return
	}

	for i, u := range used {
		add(result, append(cur, u), used[i+1:],k)
	}
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-29/007701.png)
