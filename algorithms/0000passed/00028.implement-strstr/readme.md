#### 题目
<p>实现&nbsp;<a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strStr()</a>&nbsp;函数。</p>

<p>给定一个&nbsp;haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回&nbsp; <strong>-1</strong>。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> haystack = &quot;hello&quot;, needle = &quot;ll&quot;
<strong>输出:</strong> 2
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> haystack = &quot;aaaaa&quot;, needle = &quot;bba&quot;
<strong>输出:</strong> -1
</pre>

<p><strong>说明:</strong></p>

<p>当&nbsp;<code>needle</code>&nbsp;是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。</p>

<p>对于本题而言，当&nbsp;<code>needle</code>&nbsp;是空字符串时我们应当返回 0 。这与C语言的&nbsp;<a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strstr()</a>&nbsp;以及 Java的&nbsp;<a href="https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)" target="_blank">indexOf()</a>&nbsp;定义相符。</p>


 #### 题解
 1、暴力法
 移动 haystack 的坐标值，比较坐标值和 needle的值。
 ```go
 func strStr(haystack string, needle string) int {
 	for i := 0; i <= len(haystack)-len(needle); i++ {
 		var j int
 		for j = 0; j < len(needle); j++ {
 			if haystack[i+j] != needle[j] {
 				break
 			}
 		}
 		if j == len(needle) {
 			return i
 		}
 	}
 	return -1
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002801.png)
时间复杂度O(n^2^)，空间复杂度O(1)
 
 2、子串比较
 其实在haystack字符上移动时，可以获取当前字符的len(needle)长度与needle进行比较
 ```go
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-17/002802.png)
 时间复杂度O(n^2^)，以为字符串比较的时间复杂度是O(n)，空间复杂度O(1)