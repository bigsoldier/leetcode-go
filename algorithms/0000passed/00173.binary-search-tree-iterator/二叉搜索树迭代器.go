package code

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{[]*TreeNode{}}
	it.push(root)
	return it
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.push(top.Right)
	return top.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func (this *BSTIterator) push(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
