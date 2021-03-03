package code

type NumMatrix struct {
	arr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	for i, row := range matrix {
		sum[i] = make([]int, len(row)+1)
		for j, v := range row {
			sum[i][j+1] = sum[i][j] + v
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (ans int) {
	for i := row1; i <= row2; i++ {
		ans += this.arr[i][col2+1] - this.arr[i][col1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
