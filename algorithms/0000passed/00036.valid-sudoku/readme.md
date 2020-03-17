#### 题目
<p>判断一个&nbsp;9x9 的数独是否有效。只需要<strong>根据以下规则</strong>，验证已经填入的数字是否有效即可。</p>

<ol>
	<li>数字&nbsp;<code>1-9</code>&nbsp;在每一行只能出现一次。</li>
	<li>数字&nbsp;<code>1-9</code>&nbsp;在每一列只能出现一次。</li>
	<li>数字&nbsp;<code>1-9</code>&nbsp;在每一个以粗实线分隔的&nbsp;<code>3x3</code>&nbsp;宫内只能出现一次。</li>
</ol>

<p><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png" style="height: 250px; width: 250px;"></p>

<p><small>上图是一个部分填充的有效的数独。</small></p>

<p>数独部分空格内已填入了数字，空白格用&nbsp;<code>&#39;.&#39;</code>&nbsp;表示。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong>
[
  [&quot;5&quot;,&quot;3&quot;,&quot;.&quot;,&quot;.&quot;,&quot;7&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;],
  [&quot;6&quot;,&quot;.&quot;,&quot;.&quot;,&quot;1&quot;,&quot;9&quot;,&quot;5&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;],
  [&quot;.&quot;,&quot;9&quot;,&quot;8&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;6&quot;,&quot;.&quot;],
  [&quot;8&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;6&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;3&quot;],
  [&quot;4&quot;,&quot;.&quot;,&quot;.&quot;,&quot;8&quot;,&quot;.&quot;,&quot;3&quot;,&quot;.&quot;,&quot;.&quot;,&quot;1&quot;],
  [&quot;7&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;2&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;6&quot;],
  [&quot;.&quot;,&quot;6&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;2&quot;,&quot;8&quot;,&quot;.&quot;],
  [&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;4&quot;,&quot;1&quot;,&quot;9&quot;,&quot;.&quot;,&quot;.&quot;,&quot;5&quot;],
  [&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;8&quot;,&quot;.&quot;,&quot;.&quot;,&quot;7&quot;,&quot;9&quot;]
]
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong>
[
&nbsp; [&quot;8&quot;,&quot;3&quot;,&quot;.&quot;,&quot;.&quot;,&quot;7&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;],
&nbsp; [&quot;6&quot;,&quot;.&quot;,&quot;.&quot;,&quot;1&quot;,&quot;9&quot;,&quot;5&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;],
&nbsp; [&quot;.&quot;,&quot;9&quot;,&quot;8&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;6&quot;,&quot;.&quot;],
&nbsp; [&quot;8&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;6&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;3&quot;],
&nbsp; [&quot;4&quot;,&quot;.&quot;,&quot;.&quot;,&quot;8&quot;,&quot;.&quot;,&quot;3&quot;,&quot;.&quot;,&quot;.&quot;,&quot;1&quot;],
&nbsp; [&quot;7&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;2&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;6&quot;],
&nbsp; [&quot;.&quot;,&quot;6&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;2&quot;,&quot;8&quot;,&quot;.&quot;],
&nbsp; [&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;4&quot;,&quot;1&quot;,&quot;9&quot;,&quot;.&quot;,&quot;.&quot;,&quot;5&quot;],
&nbsp; [&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;.&quot;,&quot;8&quot;,&quot;.&quot;,&quot;.&quot;,&quot;7&quot;,&quot;9&quot;]
]
<strong>输出:</strong> false
<strong>解释:</strong> 除了第一行的第一个数字从<strong> 5</strong> 改为 <strong>8 </strong>以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>一个有效的数独（部分已被填充）不一定是可解的。</li>
	<li>只需要根据以上规则，验证已经填入的数字是否有效即可。</li>
	<li>给定数独序列只包含数字&nbsp;<code>1-9</code>&nbsp;和字符&nbsp;<code>&#39;.&#39;</code>&nbsp;。</li>
	<li>给定数独永远是&nbsp;<code>9x9</code>&nbsp;形式的。</li>
</ul>


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