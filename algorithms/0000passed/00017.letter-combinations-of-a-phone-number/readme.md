#### 题目
<p>给定一个仅包含数字&nbsp;<code>2-9</code>&nbsp;的字符串，返回所有它能表示的字母组合。</p>

<p>给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png" style="width: 200px;"></p>

<p><strong>示例:</strong></p>

<pre><strong>输入：</strong>&quot;23&quot;
<strong>输出：</strong>[&quot;ad&quot;, &quot;ae&quot;, &quot;af&quot;, &quot;bd&quot;, &quot;be&quot;, &quot;bf&quot;, &quot;cd&quot;, &quot;ce&quot;, &quot;cf&quot;].
</pre>

<p><strong>说明:</strong><br>
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。</p>


 #### 题解
 这道题本来用循环去做，但是发现循环的个数不确定，所以得使用递归
 ```go
 var (
 	letterMap = []string{
 	"",
 	"",
 	"abc",
 	"def",
 	"ghi",
 	"jkl",
 	"mno",
 	"pqrs",
 	"tuv",
 	"wxyz",
 }
 	result []string
 )
 
 func letterCombinations(digits string) []string {
 	if len(digits) == 0 {
 		return nil
 	}
 	getString(digits,"")
 	return result
 }
 
 func getString(digits, s string) {
 	if len(digits) == 0 {
 		result = append(result,s) // 最后取所有的字母
 		return
 	}
 
 	letter := letterMap[digits[0]-'0']
 	for i:=0;i<len(letter);i++ {
 		getString(digits[1:],s+string(letter[i]))
 	}
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-13/001702.png)
 时间复杂度O(3^m^*4^n^)，空间复杂度O(len(3*m+4*n)),m和n表示数字对应的字符