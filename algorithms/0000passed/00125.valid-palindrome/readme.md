#### 题目
<p>给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。</p>

<p><strong>说明：</strong>本题中，我们将空字符串定义为有效的回文串。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> &quot;A man, a plan, a canal: Panama&quot;
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> &quot;race a car&quot;
<strong>输出:</strong> false
</pre>


 #### 题解
 ```go
 func isPalindrome(s string) bool {
 	if len(s) == 0 {
 		return true
 	}
 	var ret string
 	for i := 0; i < len(s); i++ {
 		if isAlnum(s[i]) {
 			ret += string(s[i])
 		}
 	}
 	ret = strings.ToLower(ret)
 	for i := 0; i < len(ret)/2; i++ {
 		if ret[i] != ret[len(ret)-i-1] {
 			return false
 		}
 	}
 	return true
 }
 
 func isAlnum(ch byte) bool {
 	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
 }
```
 时间复杂度O(n),空间复杂度O(n)
 
 ```go
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var left,right = 0,len(s)-1
	for left < right {
		for left < right && !isAlnum(s[left]) {
			left++
		}
		for left < right && !isAlnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isAlnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
```
 时间复杂度O(n),空间复杂度O(1)