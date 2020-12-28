#### 题目
<p>给定两个单词（<em>beginWord&nbsp;</em>和 <em>endWord</em>）和一个字典，找到从&nbsp;<em>beginWord</em> 到&nbsp;<em>endWord</em> 的最短转换序列的长度。转换需遵循如下规则：</p>

<ol>
	<li>每次转换只能改变一个字母。</li>
	<li>转换过程中的中间单词必须是字典中的单词。</li>
</ol>

<p><strong>说明:</strong></p>

<ul>
	<li>如果不存在这样的转换序列，返回 0。</li>
	<li>所有单词具有相同的长度。</li>
	<li>所有单词只由小写字母组成。</li>
	<li>字典中不存在重复的单词。</li>
	<li>你可以假设 <em>beginWord</em> 和 <em>endWord </em>是非空的，且二者不相同。</li>
</ul>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong>
beginWord = &quot;hit&quot;,
endWord = &quot;cog&quot;,
wordList = [&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;lot&quot;,&quot;log&quot;,&quot;cog&quot;]

<strong>输出: </strong>5

<strong>解释: </strong>一个最短转换序列是 &quot;hit&quot; -&gt; &quot;hot&quot; -&gt; &quot;dot&quot; -&gt; &quot;dog&quot; -&gt; &quot;cog&quot;,
     返回它的长度 5。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>
beginWord = &quot;hit&quot;
endWord = &quot;cog&quot;
wordList = [&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;lot&quot;,&quot;log&quot;]

<strong>输出:</strong>&nbsp;0

<strong>解释:</strong>&nbsp;<em>endWord</em> &quot;cog&quot; 不在字典中，所以无法进行转换。</pre>


 #### 题解
 广度搜索
 ```go
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}
	queue := []string{beginWord}
	level := 1
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			for c := 0; c < len(word); c++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := word[:c] + string(j) + word[c+1:]
					if wordMap[newWord] == true {
						queue = append(queue, newWord)
						wordMap[newWord] = false
					}
				}
			}
		}
		level++
	}
	return 0
}
```
 时间复杂度O(n*w^2^),空间复杂度O(n*w^2^)
 
 广度搜索+无向图
 ```go
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}
```
 时间复杂度O(n*w^2^),空间复杂度O(n*w^2^)
 
 双向搜索
 ```go
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	distBegin := make([]int, len(wordId))
	for i := range distBegin {
		distBegin[i] = inf
	}
	distBegin[beginId] = 0
	queueBegin := []int{beginId}
	
	distEnd := make([]int,len(wordId))
	for i := range distEnd {
		distEnd[i] = inf
	}
	distEnd[endId] = 0
	queueEnd := []int{endId}
	
	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		q := queueBegin
		queueBegin = nil
		for _, v := range q {
			if distEnd[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distBegin[w] == inf {
					distBegin[w] = distBegin[v] +1
					queueBegin = append(queueBegin, w)
				}
			}
		}
		
		q = queueEnd
		queueEnd = nil
		for _, v := range q {
			if distBegin[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distEnd[w] == inf {
					distEnd[w] = distEnd[v] +1
					queueEnd = append(queueEnd, w)
				}
			}
		}
	}
	return 0
}
```
时间复杂度O(n*w^2^),空间复杂度O(n*w^2^)