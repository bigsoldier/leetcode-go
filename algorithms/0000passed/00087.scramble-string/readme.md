#### 题目
<p>给定一个字符串&nbsp;<em>s1</em>，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。</p>

<p>下图是字符串&nbsp;<em>s1</em>&nbsp;=&nbsp;<code>&quot;great&quot;</code>&nbsp;的一种可能的表示形式。</p>

<pre>    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
</pre>

<p>在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。</p>

<p>例如，如果我们挑选非叶节点&nbsp;<code>&quot;gr&quot;</code>&nbsp;，交换它的两个子节点，将会产生扰乱字符串&nbsp;<code>&quot;rgeat&quot;</code>&nbsp;。</p>

<pre>    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
</pre>

<p>我们将&nbsp;<code>&quot;rgeat&rdquo;</code>&nbsp;称作&nbsp;<code>&quot;great&quot;</code>&nbsp;的一个扰乱字符串。</p>

<p>同样地，如果我们继续交换节点&nbsp;<code>&quot;eat&quot;</code>&nbsp;和&nbsp;<code>&quot;at&quot;</code>&nbsp;的子节点，将会产生另一个新的扰乱字符串&nbsp;<code>&quot;rgtae&quot;</code>&nbsp;。</p>

<pre>    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
</pre>

<p>我们将&nbsp;<code>&quot;rgtae&rdquo;</code>&nbsp;称作&nbsp;<code>&quot;great&quot;</code>&nbsp;的一个扰乱字符串。</p>

<p>给出两个长度相等的字符串 <em>s1 </em>和&nbsp;<em>s2</em>，判断&nbsp;<em>s2&nbsp;</em>是否是&nbsp;<em>s1&nbsp;</em>的扰乱字符串。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> s1 = &quot;great&quot;, s2 = &quot;rgeat&quot;
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> s1 = &quot;abcde&quot;, s2 = &quot;caebd&quot;
<strong>输出:</strong> false</pre>


 #### 题解
 递归：遍历所有可能情况，分割点从1到n-1，s1分成[0,i)[i,n],对应的s2可能为[0,i)[i,n]或[0,n-i)[n-i,n]
 ```go
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	// 检查字符是否相等
	n := len(s1)
	var chars = make([]int,256)
	for i := 0; i < n; i++ {
		chars[s1[i]] ++
		chars[s2[i]] --
	}
	for _, char := range chars {
		if char != 0 {
			return false
		}
	}

	// 检查字符串是否scramble
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-27/008701.png)
 