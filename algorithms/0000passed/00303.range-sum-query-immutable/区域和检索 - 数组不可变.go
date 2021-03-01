package code

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	var sum = make([]int, len(nums)+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	return NumArray{arr: sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.arr[j+1] - this.arr[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
