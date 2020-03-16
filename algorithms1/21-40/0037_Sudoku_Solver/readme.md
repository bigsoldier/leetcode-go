#### 题目
![题目](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-24/003701.png)

#### 题解
1、暴力法
将每种可能性都验证一遍，一个格子有9种可能，共81个格子，时间复杂度最多9^81^

2、回溯法
依次往空格中添加 *1-9* ，保证满足条件
```go
func solveSudoku(board [][]byte) {
	solveSudo(board,0)
}

func solveSudo(board [][]byte, index int) bool {
	if index == 81 {
		return true
	}

	row,col := index/9,index%9
	if board[row][col] != '.' {
		return solveSudo(board,index+1)
	}

	boxI,boxJ := (row/3)*3,(col/3)*3

	isValid := func(b byte) bool { // 校验横行、纵行、9宫格的数字是否符合标准
		for i := 0; i < 9; i++ {
			if board[row][i] == b || board[i][col] == b || board[boxI+i/3][boxJ+i%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		// 这里尝试填数字
		if isValid(b) { // 如果有符合的数字，将该空赋值并校验下一个空
			board[row][col] = b
			if solveSudo(board,index+1) {
				return true
			}
		}
	}

	// 沒有找到，置'.'
	board[row][col] = '.'

	return false
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-11/003702.png)