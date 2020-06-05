#### 题目
<p>给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> &quot;25525511135&quot;
<strong>输出:</strong> <code>[&quot;255.255.11.135&quot;, &quot;255.255.111.35&quot;]</code></pre>


 #### 题解
 递归法
 ```go
func restoreIpAddresses(s string) []string {
	var result []string
	restoreIpAddress(nil,s,&result,0)
	return result
}

func restoreIpAddress(ans []string,s string, result *[]string, depth int) {
	if depth == 4 {
		if len(s) != 0 {
			return
		}
		*result = append(*result, strings.Join(ans,"."))
		return
	}
	for i := 1; i < 4; i++ {
		if i <= len(s) && valid(s[:i]) {
			restoreIpAddress(append(ans, s[:i]),s[i:],result,depth+1)
		}
	}
}

// 判断字符串是否在0-255之间
func valid(s string) bool {
	switch len(s) {
	case 1, 2:
		if s[0] == '0' && len(s) != 1{
			return false
		}
		return true
	case 3:
		if s[0] == '0' && len(s) != 1{
			return false
		}
		if s < "256" {
			return true
		}
	}
	return false
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-04/009301.png)
 