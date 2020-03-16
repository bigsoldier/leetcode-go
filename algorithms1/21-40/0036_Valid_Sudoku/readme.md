#### 题目
![题目](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-24/003601.png)

#### 题解
使用三个二维map存储
```go
func isValidSudoku(board [][]byte) bool {
	var rows,colums,boxes = make(map[int]map[byte]int),make(map[int]map[byte]int),make(map[int]map[byte]int)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]int)
		colums[i] = make(map[byte]int)
		boxes[i] = make(map[byte]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			bt := board[i][j]
			if bt != '.' {
				_,ok := rows[i][bt]
				if ok {
					return false
				}
				rows[i][bt] = j

				_,ok = colums[j][bt]
				if ok {
					return false
				}
				colums[j][bt] = i

				_,ok = boxes[i/3 + j/3 * 3][bt]
				if ok {
					return false
				}
				boxes[i/3 + j/3 * 3][bt] = i
			}
		}
	}
	return true
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-10/003602.png)
时间复杂度O(n^2^)，空间复杂度O(n^2^)

**优化**
不管横竖还是格子，都是要统计9个的，使用数组作为存储，减少内存分配
```go
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidSudokuRow(board,i) {
			return false
		}
		if !isValidSudokuCol(board, i) {
			return false
		}
		if !isValidSudokuBox(board,i) {
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var rows  [10]bool
	for i := 0; i < 9; i++ {
		num := convertNum(board[row][i])
		if num < 0 {
			continue
		}
		if rows[num] {
			return false
		}
		rows[num] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var rows  [10]bool
	for i := 0; i < 9; i++ {
		num := convertNum(board[i][col])
		if num < 0 {
			continue
		}
		if rows[num] {
			return false
		}
		rows[num] = true
	}
	return true
}

func isValidSudokuBox(board [][]byte, box int) bool {
	var rows [10]bool

	row := (box / 3)*3
	col := (box%3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := convertNum(board[i+row][j+col])
			if num < 0 {
				continue
			}
			if rows[num] {
				return false
			}
			rows[num] = true
		}
	}
	return true
}

func convertNum(bt byte) int {
	if bt == '.' {
		return -1
	}
	return int(bt-'0')
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-10/003603.png)
时间复杂度O(n^2^),空间复杂度O(n)