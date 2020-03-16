#### 题目

将一个给定的字符串根据给定的行数，以从上到下、从左到右进行 Z 字形排序。

比如输入字符串为 *“LEETCODEISHIRING”* 行数为3，排列如下：

```
L	C	I	R
E T O E S I	I G
E	D	H	N
```

之后，需要从左到右逐行读取，结果为 “LCIRETOESIIGEDHN”

**示例1：**

```
输入：s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
```

**示例2：**

```
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
```



#### 题解
1、找规律
*Z* 字形很容易找到规律，假如 *numRows=3* ，那么周期为4；假如 *numRows=4* ，那么周期为6。

那么第一行的字符索引为 *2 numRows -2*

最后一行的字符索引为 *（2 numRows -2）+ numRows-1*

内部的行 *i* 的字符索引为 *(2 numRows -2)+i* 和 *(2 numRows -2)-i*

```go
func convert(s string, numRows int) string {
	if 1 == numRows {
		return s
	}

	var ret string
	var T = 2 * numRows - 2 // 周期
	for i := 0; i < numRows; i++ {
		for j := 0; i + j < len(s); j += T {
			ret += s[i+j:i+j+1]
			if i != 0 && i != numRows-1 && j+T-i < len(s) {
				ret += s[j+T-i:j+T-i+1]
			}
		}
	}
	return string(ret)
}
```

![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-14/000601.png)
时间复杂度：O(n)
空间复杂度：O(n)

2、按行排序
建立一个行列表存放每行的数据，遍历原字符串，观察可以得出，移到最上面和最下面的时候，当前方向才会改变。
0     6      12
1   5 7   11 13
2 4   8 10   14
3     9      15
```go
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret = make([]string,numRows)
	var T = 2 * numRows - 2 // 周期（必定为偶数）
	for index, bt := range []byte(s) {
		rows := index % T
		if rows > T/2 {
			ret[T-rows] += string(bt)
		} else {
			ret[rows] += string(bt)
		}
	}
	return strings.Join(ret,"")
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-15/000602.png)
时间复杂度：O(n)
空间复杂度：O(n)