#### 题目
给定两个整数，被整数 *dividend* 和除数 *divisor* 。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 *dividend* 除以除数 *divisor* 得到的商。
**示例1：**
```
输入：dividend = 10，divisor = 3
输出：3
```
**示例2：**
```
输入：dividend = 7，divisor = -3
输出：-2
```
**说明**
- 被除数和除数均为32位有符号整数
- 除数不为0
- 假设我们环境只能存储32位有符号整数，其数值范围是 [-2^31^,2^31-1]。本题中，如果除法结果溢出，则返回 2^31-1。

#### 题解
1、全都转化为正数计算
```go
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32	// 是最小的那个数
	}

	signal := 1 	// 正符号
	if (dividend > 0 && divisor < 0) || (dividend < 0 &&divisor > 0) {
		signal = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	result := div(dividend,divisor)
	if signal > 0 {
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		return result
	}
	return -result
}

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}
	var count = 1
	tb := divisor
	for tb + tb <= dividend {
		tb += tb
		count += count
	}
	return count + div(dividend - tb,divisor)
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-19/002901.png)

2、全都转化为负数计算
正数需要考虑边界问题，使用负数就不用考虑边界问题了。
```go
func divide(dividend int, divisor int) int {
	signal := -1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		signal = 1
	}
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	var result int
	for dividend <= divisor {
		var tmpDivisor = divisor
		var tmpResult = -1
		for dividend <= (tmpDivisor << 1) {
			if tmpDivisor <= (math.MinInt32>>1) {
				break
			}
			tmpDivisor = tmpDivisor<<1
			tmpResult = tmpResult<<1
		}
		dividend -= tmpDivisor
		result += tmpResult
	}
	if signal < 0 {
		if result <= math.MinInt32 {
			return math.MaxInt32
		}
		result = -result
	}
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-19/002902.png)