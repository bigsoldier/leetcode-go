#### 题目
罗马数字包含以下七种字符：I,V,X,L,C,D,M
```$xslt
字符      数值
I           1
V           5
X           10
L           50
C           100
D           500
M           1000
```
例如，罗马数字2写作II，12写作XII，27写作XXVII。

通常情况下，罗马数字中小的数字在大的数字右边。但存在特例，例如4写作IV。这种特殊规则适用于以下6种情况：
- I 可以放在 V(5) 和 X(10) 的左边，来表示 4 和 9
- X 可以放在 L(50) 和 C(100) 的左边，来表示 40 和 90
- C 可以放在 D(500) 和 M(1000) 的左边，来表示 400 和 900

给定一个整数，将其转为罗马数字，输入确保在1到3999的范围内。

#### 题解
其实只需要比较前一位和后一位的大小即可，如果前一位小于后一位，减去即可。（IV即5-1=4）
```go
func romanToInt(s string) int {
	var hashNum = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	var preNum = hashNum[s[0]]
	for i := 1; i < len(s); i++ {
		num := hashNum[s[i]]
		if preNum < num {
			result -= preNum
		} else {
			result += preNum
		}
		preNum = num
	}
	result += preNum
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-06/001301.png)
时间复杂度O(n),空间复杂度O(1)

**优化**
将map转成switch，减少空间消耗
```go
func romanToInt(s string) int {
	var result int
	var preNum = getValue(s[0])
	for i := 1; i < len(s); i++ {
		num := getValue(s[i])
		if preNum < num {
			result -= preNum
		} else {
			result += preNum
		}
		preNum = num
	}
	result += preNum
	return result
}

func getValue(bt byte) int {
	switch bt {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-06/001302.png)
时间复杂度O(n),空间复杂度O(1)