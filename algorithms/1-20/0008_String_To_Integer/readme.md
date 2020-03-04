#### 题目
请你来实现一个 *atoi* 函数，使其能将字符串转成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找第一个非空格的字符为止。
当我们找到的第一个非空字符为正或负号时，则将该符号与之后面尽可能多的**连续**数字组合起来。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−2^31^,  2^31^ − 1]。如果数值超过这个范围，请返回  INT_MAX (2^31^ − 1) 或 INT_MIN (−2^31^) 。

**示例1：**
```
输入：“42”
输出：42
```
**示例2：**
```
输入：“-42”
输出：-42
解释：第一个非空白字符为‘-’，是负号
```
**示例3：**
```
输入：“4139 with words”
输出：4139
解释：转换截止于数字‘3’，因为下一个字符不是数字
```
**示例4：**
```
输入：“words and 987”
输出：0
解释：第一个非空字符是‘w'，但不是数字或正、负号。
```
**示例5：**
```
输入：“-91283472332”
输出：-2147483648
解释：数组超过32位有符数字，返回32位最小数字
```

#### 题解
使用正则表达式取出数字和正负号。
```go
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	regx := regexp.MustCompile(`^[+-]?\d*`)
	ret := regx.FindString(str)
	number,_ := strconv.Atoi(ret)
	if number > math.MaxInt32 {
		return math.MaxInt32
	} else if number < math.MinInt32 {
		return math.MinInt32
	} else {
		return number
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-16/000801.png)