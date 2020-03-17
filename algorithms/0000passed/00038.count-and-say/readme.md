#### 题目
<p>「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：</p>

<pre>1.     1
2.     11
3.     21
4.     1211
5.     111221
</pre>

<p><code>1</code>&nbsp;被读作&nbsp;&nbsp;<code>&quot;one 1&quot;</code>&nbsp;&nbsp;(<code>&quot;一个一&quot;</code>) , 即&nbsp;<code>11</code>。<br>
<code>11</code> 被读作&nbsp;<code>&quot;two 1s&quot;</code>&nbsp;(<code>&quot;两个一&quot;</code>）, 即&nbsp;<code>21</code>。<br>
<code>21</code> 被读作&nbsp;<code>&quot;one 2&quot;</code>, &nbsp;&quot;<code>one 1&quot;</code>&nbsp;（<code>&quot;一个二&quot;</code>&nbsp;,&nbsp;&nbsp;<code>&quot;一个一&quot;</code>)&nbsp;, 即&nbsp;<code>1211</code>。</p>

<p>给定一个正整数 <em>n</em>（1 &le;&nbsp;<em>n</em>&nbsp;&le; 30），输出外观数列的第 <em>n</em> 项。</p>

<p>注意：整数序列中的每一项将表示为一个字符串。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 1
<strong>输出:</strong> &quot;1&quot;
<strong>解释：</strong>这是一个基本样例。</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> 4
<strong>输出:</strong> &quot;1211&quot;
<strong>解释：</strong>当 n = 3 时，序列是 &quot;21&quot;，其中我们有 &quot;2&quot; 和 &quot;1&quot; 两组，&quot;2&quot; 可以读作 &quot;12&quot;，也就是出现频次 = 1 而 值 = 2；类似 &quot;1&quot; 可以读作 &quot;11&quot;。所以答案是 &quot;12&quot; 和 &quot;11&quot; 组合在一起，也就是 &quot;1211&quot;。</pre>


 #### 题解
 递归
 ```go
 func countAndSay(n int) string {
 	if n == 1 {
 		return "1"
 	}
 	ret := countAndSay(n-1)
 
 	var result string
 	var count = 1
 	for i := 1; i < len(ret); i++ {
 		if ret[i] == ret[i-1] {
 			count++
 		} else {
 			result += strconv.Itoa(count) + ret[i-1:i]
 			count = 1
 		}
 	}
 	result += strconv.Itoa(count) + ret[len(ret)-1:]
 	return result
 }
 
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-11/003802.png)
 
 两重循环
 ```go
 func countAndSay(n int) string {
 	buf := []byte{'1'}
 
 	for n > 1 {
 		buf = say(buf)
 		n--
 	}
 
 	return string(buf)
 }
 
 func say(buf []byte) []byte {
 	// res 长度不会超过 buf 的两倍，所以，可以事先指定容量，加快append的速度
 	res := make([]byte, 0, len(buf)*2)
 
 	i, j := 0, 1
 	for i < len(buf) {
 		// 利用 j ，找到下一个不同的元素
 		for j < len(buf) && buf[j] == buf[i] {
 			j++
 		}
 
 		// res 中 res[i] 表示 res[i+1] 的个数，i 为0,2,4,6,...
 		res = append(res, byte(j-i+'0'), buf[i])
 
 		// 移动 i 到 j
 		i = j
 	}
 
 	return res
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-11/003803.png)