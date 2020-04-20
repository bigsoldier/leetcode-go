#### 题目
<p>给定一个仅包含大小写字母和空格&nbsp;<code>&#39; &#39;</code>&nbsp;的字符串 <code>s</code>，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。</p>

<p>如果不存在最后一个单词，请返回 0&nbsp;。</p>

<p><strong>说明：</strong>一个单词是指仅由字母组成、不包含任何空格字符的 <strong>最大子字符串</strong>。</p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> &quot;Hello World&quot;
<strong>输出:</strong> 5
</pre>


 #### 题解
 先去除字符串右边可能存在的空格，然后再从右向左遍历找到空格，计数
 ```go
func lengthOfLastWord(s string) int {
	// 去除右边的空格
	//s = strings.TrimRight(s," ") // 等价于下面的代码
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			s = s[:i+1]
			break
		}
	}
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}
	return count
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-16/005801.png)
时间复杂度O(n),空间复杂度O(1)