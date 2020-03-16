#### 题目
给定 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n=3，生成结果为：
```
[
    "((()))",
    "(())()",
    "(()())",
    "()(())",
    "()()()"    
]
```

#### 题解
1、暴力法
列举出所有的括号可能性，然后校验括号是否有效
时间复杂度O(2^2n^*n)，每个位置有两种可能，每个都需要O(n)校验
空间复杂度O(2^2n^*n)

2、回溯法
优化暴力法，计数左右括号，使得右括号不大于左括号
```go
var result []string

func generateParenthesis(n int) []string {
	result = []string{}
	backTrace("",0,0,n)
	return result
}

func backTrace(cursor string, left, right, max int) {
	if len(cursor) == max*2 { // 生成返回
		result = append(result, cursor)
		return
	}

	if left < max {
		backTrace(cursor+"(",left+1,right,max)
	}
	if left > right {
		backTrace(cursor+")",left,right+1,max)
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-16/002201.png)

3、动态规划
首位一定是左括号“（”，那么一定有一个与之匹配的右括号“）”，可以用表达式
“（”+i=p所有的括号排列组合+“）”+i=q所有的括号排列组合，
其中 *p+q=n-1*，且p q均非负整数。
```go
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var result = make([][]string,n+1)
	result[0] = []string{""}		// 0组括号的情况
	result[1] = []string{"()"}		// 1组括号的情况
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {	// 遍历p q，其中p+q=i-1
			for _, k1 := range result[j] {
				for _, k2 := range result[i-1-j] {
					result[i] = append(result[i], "("+string(k1)+")"+string(k2))
				}
			}
		}
	}
	return result[n]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-16/002202.png)
时间复杂度O(n^4^)