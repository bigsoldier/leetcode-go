package code

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
			dfs(board, trie, i, j, &ans)
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
		dfs(board, trie, x-1, y, result)
	}
	if x < len(board)-1 {
		dfs(board, trie, x+1, y, result)
	}
	if y > 0 {
		dfs(board, trie, x, y-1, result)
	}
	if y < len(board[0])-1 {
		dfs(board, trie, x, y+1, result)
	}

	board[x][y] = c
}

// 前缀树
type Trie struct {
	word     string
	children [26]*Trie
}

// 构造前缀树
func buildTrie(words []string) *Trie {
	root := new(Trie)
	for _, word := range words {
		curr := root
		for _, c := range word {
			idx := c - 'a'
			if curr.children[idx] == nil {
				curr.children[idx] = new(Trie)
			}
			curr = curr.children[idx]
		}
		curr.word = word
	}
	return root
}
