#### 题目
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-13/001701.png)

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