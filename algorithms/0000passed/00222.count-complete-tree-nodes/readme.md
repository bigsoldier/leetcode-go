#### 题目
<p>给出一个<strong>完全二叉树</strong>，求出该树的节点个数。</p>

<p><strong>说明：</strong></p>

<p><a href="https://baike.baidu.com/item/%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91/7773232?fr=aladdin">完全二叉树</a>的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~&nbsp;2<sup>h</sup>&nbsp;个节点。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 
    1
   / \
  2   3
 / \  /
4  5 6

<strong>输出:</strong> 6</pre>


 #### 题解
 由于是完全二叉树，所以节点个数[2^h^,2^h+1^-1]
 ```go
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var level int
	for node := root; node.Left != nil ; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1<<(level-1)
		node := root
		for node != nil && bits > 0 {
            // 父节点为i，左孩子是2*i+1,右孩子是2*i+2
			if bits&k == 0 { 
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>=1
		}
		return node==nil
	})-1
}
```
 时间复杂度O((logn)^2^),空间复杂度O(1)