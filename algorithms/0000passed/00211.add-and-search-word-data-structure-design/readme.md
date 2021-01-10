#### 题目
<p>设计一个支持以下两种操作的数据结构：</p>

<pre>void addWord(word)
bool search(word)
</pre>

<p>search(word)&nbsp;可以搜索文字或正则表达式字符串，字符串只包含字母&nbsp;<code>.</code>&nbsp;或&nbsp;<code>a-z</code>&nbsp;。&nbsp;<code>.</code> 可以表示任何一个字母。</p>

<p><strong>示例:</strong></p>

<pre>addWord(&quot;bad&quot;)
addWord(&quot;dad&quot;)
addWord(&quot;mad&quot;)
search(&quot;pad&quot;) -&gt; false
search(&quot;bad&quot;) -&gt; true
search(&quot;.ad&quot;) -&gt; true
search(&quot;b..&quot;) -&gt; true
</pre>

<p><strong>说明:</strong></p>

<p>你可以假设所有单词都是由小写字母 <code>a-z</code>&nbsp;组成的。</p>


 #### 题解
 同208
 ```go
type WordDictionary struct {
	isWord bool
	children map[rune]*WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{children: make(map[rune]*WordDictionary)}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	parent := this
	for _, ch := range word {
		if child,ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &WordDictionary{children: make(map[rune]*WordDictionary)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	parent := this
	for i, ch := range word {
		if ch == '.' {
			isMatch := false
			for _, v := range parent.children {
				if v.Search(word[i+1:]) {
					isMatch = true
				}
			}
			return isMatch
		} else if _, ok := parent.children[ch]; !ok {
			return false
		}
		parent = parent.children[ch]
	}
	return len(parent.children)==0 || parent.isWord
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
```