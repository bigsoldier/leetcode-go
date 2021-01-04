#### 题目
<p>给定一个正整数，返回它在 Excel 表中相对应的列名称。</p>

<p>例如，</p>

<pre>    1 -&gt; A
    2 -&gt; B
    3 -&gt; C
    ...
    26 -&gt; Z
    27 -&gt; AA
    28 -&gt; AB 
    ...
</pre>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 1
<strong>输出:</strong> &quot;A&quot;
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 28
<strong>输出:</strong> &quot;AB&quot;
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong> 701
<strong>输出:</strong> &quot;ZY&quot;
</pre>


 #### 题解
 类似于 26 进制
 ```go
func convertToTitle(n int) (ans string) {
	for n > 0 {
		n--
		ans = string('A'+n%26) + ans
		n = n/26
	}
	return
}
```
 时间复杂度O(n),空间复杂度O(1)