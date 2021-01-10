#### 题目
<p>实现一个 Trie (前缀树)，包含&nbsp;<code>insert</code>,&nbsp;<code>search</code>, 和&nbsp;<code>startsWith</code>&nbsp;这三个操作。</p>

<p><strong>示例:</strong></p>

<pre>Trie trie = new Trie();

trie.insert(&quot;apple&quot;);
trie.search(&quot;apple&quot;);   // 返回 true
trie.search(&quot;app&quot;);     // 返回 false
trie.startsWith(&quot;app&quot;); // 返回 true
trie.insert(&quot;app&quot;);   
trie.search(&quot;app&quot;);     // 返回 true</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>你可以假设所有的输入都是由小写字母&nbsp;<code>a-z</code>&nbsp;构成的。</li>
	<li>保证所有输入均为非空字符串。</li>
</ul>


 #### 题解
 字母映射表，多叉树
 ```go
type Trie struct {
	isEnd bool // 该节点是否是一个串的结束
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
```