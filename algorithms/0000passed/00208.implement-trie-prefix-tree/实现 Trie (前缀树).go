package code

type Trie struct {
	isEnd bool           // 该节点是否是一个串的结束
	child map[rune]*Trie // 字母映射表
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{child: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.child[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{child: make(map[rune]*Trie)}
			parent.child[ch] = newChild
			parent = newChild
		}
	}
	parent.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.child[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.child[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
