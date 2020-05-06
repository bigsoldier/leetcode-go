#### 题目
<p>给定一个二维网格和一个单词，找出该单词是否存在于网格中。</p>

<p>单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中&ldquo;相邻&rdquo;单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。</p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre>board =
[
  [&#39;A&#39;,&#39;B&#39;,&#39;C&#39;,&#39;E&#39;],
  [&#39;S&#39;,&#39;F&#39;,&#39;C&#39;,&#39;S&#39;],
  [&#39;A&#39;,&#39;D&#39;,&#39;E&#39;,&#39;E&#39;]
]

给定 word = &quot;<strong>ABCCED</strong>&quot;, 返回 <strong>true</strong>
给定 word = &quot;<strong>SEE</strong>&quot;, 返回 <strong>true</strong>
给定 word = &quot;<strong>ABCB</strong>&quot;, 返回 <strong>false</strong></pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>board</code> 和 <code>word</code> 中只包含大写和小写英文字母。</li>
	<li><code>1 &lt;= board.length &lt;= 200</code></li>
	<li><code>1 &lt;= board[i].length &lt;= 200</code></li>
	<li><code>1 &lt;= word.length &lt;= 10^3</code></li>
</ul>


 #### 题解
 深度搜索问题
 找到符合单词第一个字母时，需要上下左右查询第二个字母
 ```go
var (
	marked map[int]map[int]bool
	direction = [][]int{{-1,0},{0,-1},{0,1},{1,0}}
	boardM,boardN int
)

func exist(board [][]byte, word string) bool {
	boardM = len(board)
	if boardM == 0 {
		return false
	}
	boardN = len(board[0])
	marked = make(map[int]map[int]bool,boardM)
	for i := 0; i < boardM; i++ {
		marked[i] = make(map[int]bool,boardN)
	}
	for i := 0; i < boardM; i++ {
		for j := 0; j < boardN; j++ {
			if dfs(board,word,i,j,0) {
				return true
			}
		}
	}
	return false
}

// 查找(i,j)
func dfs(board [][]byte, word string, i, j, start int) bool {
	if start == len(word)-1 {
		return board[i][j] == word[start]
	}

	if board[i][j] == word[start] {
		marked[i][j] = true
		for k := 0; k < 4; k++ {
			newX := i+direction[k][0]
			newY := j+direction[k][1]
			if inArea(newX, newY) && !marked[newX][newY] {
				if dfs(board, word, newX, newY, start+1) {
					return true
				}
			}
		}
		marked[i][j] = false
	}
	return false
}

func inArea(x, y int) bool {
	return x >= 0 && x < boardM && y >= 0 && y < boardN
}

```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-06/007901.png)