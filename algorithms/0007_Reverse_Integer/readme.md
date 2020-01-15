#### 题目
给出一个 **32** 位的有符号的整数，你需要将这个整数中每位上的数字进行反转。
**示例1：**
```$xslt
输入：123
输出：321
```
**示例2：**
```$xslt
输入：-123
输出：-321
```
**示例3：**
```$xslt
输入：120
输出：21
```

#### 题解
如果不考虑溢出问题，是很简单的，利用除数余数就能解决。
增加 **32** 位有符号整数，就定了这个数最多就 31 位，因此不能大于 (2^32^-1)/10，小于 (-2^32^-1)/10。
当出现等于情况时，*2^31^-1* 的个位数为7，*-2^31^-1* 的个位数为8
```go
func reverse(x int) int {
	var temp int
	for ; x != 0; x /= 10 {
		t := x % 10
		if temp > math.MaxInt32/10 || (temp == math.MaxInt32/10 && t > 7) {
			return 0
		}
		if temp < math.MinInt32/10 || (temp == math.MinInt32/10 && t < -8) {
			return 0
		}
		temp = temp * 10 + t
	}
	return temp
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-15/000701.png)