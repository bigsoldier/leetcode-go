

#### 题解
1、找规律
*Z* 字形很容易找到规律，假如 *numRows=3* ，那么周期为4；假如 *numRows=4* ，那么周期为6
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