#### 题目
<p>给定两个字符串 <em><strong>s</strong></em> 和 <em><strong>t</strong></em>，它们只包含小写字母。</p>

<p>字符串&nbsp;<strong><em>t</em></strong>&nbsp;由字符串&nbsp;<strong><em>s</em></strong>&nbsp;随机重排，然后在随机位置添加一个字母。</p>

<p>请找出在 <em><strong>t</strong></em> 中被添加的字母。</p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre>输入：
s = &quot;abcd&quot;
t = &quot;abcde&quot;

输出：
e

解释：
&#39;e&#39; 是那个被添加的字母。
</pre>


 #### 题解
 **计数**
 因为全是小写字母，可以用数组来记录字母出现的次数
 ```go
func findTheDifference(s string, t string) byte {
	var slice [26]int
	for _, val := range s {
		slice[val-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		slice[ch-'a']--
		if slice[ch-'a'] < 0 {
			return ch
		}
	}
}
```
 时间复杂度O(n),空间复杂度O(26)
 
 **求和**
 求出s和t的ascii总值的差值
 ```go
func findTheDifference(s string, t string) byte {
	var sum = 0
	for _, v := range t {
		sum+=int(v)
	}
	for _, v := range s {
		sum-=int(v)
	}
	return byte(sum)
}
```
 时间复杂度O(n),空间复杂度O(1)

 **位运算**
 与  &: 同为1时才为1,101 & 011 = 001
 或  |: 同为0时才为0, 101 | 011 = 111
 异或 ^: 不同为1,相同为0, 101 ^ 011 = 110， 111 ^ 111 = 000
 
 根据上面的方法，将s和t所有值都进行异或，那么s和t相同的值都为0，最后剩下的值就是多出的值
 ```go
func findTheDifference(s string, t string) byte {
	var diff byte
	for i := range s {
		diff ^= s[i]^t[i]
	}
	return diff^t[len(t)-1]
}
```
 时间复杂度O(n),空间复杂度O(1)