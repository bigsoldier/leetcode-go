#### 题目
<p>给定两个单词（<em>beginWord</em> 和 <em>endWord</em>）和一个字典 <em>wordList</em>，找出所有从 <em>beginWord </em>到 <em>endWord </em>的最短转换序列。转换需遵循如下规则：</p>

<ol>
	<li>每次转换只能改变一个字母。</li>
	<li>转换过程中的中间单词必须是字典中的单词。</li>
</ol>

<p><strong>说明:</strong></p>

<ul>
	<li>如果不存在这样的转换序列，返回一个空列表。</li>
	<li>所有单词具有相同的长度。</li>
	<li>所有单词只由小写字母组成。</li>
	<li>字典中不存在重复的单词。</li>
	<li>你可以假设 <em>beginWord</em> 和 <em>endWord </em>是非空的，且二者不相同。</li>
</ul>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong>
beginWord = &quot;hit&quot;,
endWord = &quot;cog&quot;,
wordList = [&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;lot&quot;,&quot;log&quot;,&quot;cog&quot;]

<strong>输出:</strong>
[
  [&quot;hit&quot;,&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;cog&quot;],
&nbsp; [&quot;hit&quot;,&quot;hot&quot;,&quot;lot&quot;,&quot;log&quot;,&quot;cog&quot;]
]
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>
beginWord = &quot;hit&quot;
endWord = &quot;cog&quot;
wordList = [&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;lot&quot;,&quot;log&quot;]

<strong>输出: </strong>[]

<strong>解释:</strong>&nbsp;<em>endWord</em> &quot;cog&quot; 不在字典中，所以不存在符合要求的转换序列。</pre>


 #### 题解
 ```go
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}

	n := len(wordList)
	edgs := make([][]int,len(wordList))
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edgs[i] = append(edgs[i], j)
				edgs[j] = append(edgs[j], i)
			}
		}
	}

	var res [][]string
	cost := make([]int,n)
	queue := [][]int{{ids[beginWord]}}

	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else {
			for _, to := range edgs[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last]+1
					tmp := make([]int,len(now))
					copy(tmp,now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

// 比较两个字符串是否差1个相等
func transformCheck(from, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i+1:]==to[i+1:]
		}
	}
	return false
}
```
时间复杂度O(n*w^2^),空间复杂度O(n*w^2^)