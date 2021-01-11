#### 题目
<p>给定一个二维网格&nbsp;<strong>board&nbsp;</strong>和一个字典中的单词列表 <strong>words</strong>，找出所有同时在二维网格和字典中出现的单词。</p>

<p>单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中&ldquo;相邻&rdquo;单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 
<strong>words</strong> = <code>[&quot;oath&quot;,&quot;pea&quot;,&quot;eat&quot;,&quot;rain&quot;]</code> and <strong>board </strong>=
[
  [&#39;<strong>o</strong>&#39;,&#39;<strong>a</strong>&#39;,&#39;a&#39;,&#39;n&#39;],
  [&#39;e&#39;,&#39;<strong>t</strong>&#39;,&#39;<strong>a</strong>&#39;,&#39;<strong>e</strong>&#39;],
  [&#39;i&#39;,&#39;<strong>h</strong>&#39;,&#39;k&#39;,&#39;r&#39;],
  [&#39;i&#39;,&#39;f&#39;,&#39;l&#39;,&#39;v&#39;]
]

<strong>输出:&nbsp;</strong><code>[&quot;eat&quot;,&quot;oath&quot;]</code></pre>

<p><strong>说明:</strong><br>
你可以假设所有输入都由小写字母 <code>a-z</code>&nbsp;组成。</p>

<p><strong>提示:</strong></p>

<ul>
	<li>你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？</li>
	<li>如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： <a href="/problems/implement-trie-prefix-tree/description/">实现Trie（前缀树）</a>。</li>
</ul>


 #### 题解
 将要查找的单词做成前缀树
 ```go
func findWords(board [][]byte, words []string) (ans []string) {
	m := len(board)
	if m == 0 {
		return nil
	}
	n := len(board[0])
	if n == 0 {
		return nil
	}

	trie := buildTrie(words)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, trie, i, j,&ans)
		}
	}
	return
}

func dfs(board [][]byte, trie *Trie, x, y int, result *[]string) {
	c := board[x][y]
	if c == '#' || trie.children[c-'a'] == nil {
		return
	}
	trie = trie.children[c-'a']
	if len(trie.word) > 0 {
		*result = append(*result, trie.word)
		trie.word = ""
	}

	board[x][y] = '#'

	if x > 0 {
		dfs(board,trie,x-1,y,result)
	}
	if x < len(board)-1 {
		dfs(board,trie,x+1,y,result)
	}
	if y > 0 {
		dfs(board,trie,x,y-1,result)
	}
	if y < len(board[0])-1 {
		dfs(board,trie,x,y+1,result)
	}

	board[x][y] = c
}

// 前缀树
type Trie struct {
	word string
	children [26]*Trie
}

// 构造前缀树
func buildTrie(words []string) *Trie {
	root := new(Trie)
	for _, word := range words {
		curr := root
		for _, c := range word {
			idx := c-'a'
			if curr.children[idx] == nil {
				curr.children[idx] = new(Trie)
			}
			curr = curr.children[idx]
		}
		curr.word = word
	}
	return root
}
```
 时间复杂度O(M(4*3^L-1^)),空间复杂度O(n),M是二维网格的个数，L是单词的最大长度，n是字典中字母数