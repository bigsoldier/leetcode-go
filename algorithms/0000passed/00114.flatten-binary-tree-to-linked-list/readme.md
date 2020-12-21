#### 题目
<p>给定一个二叉树，<a href="https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95/8010757" target="_blank">原地</a>将它展开为链表。</p>

<p>例如，给定二叉树</p>

<pre>    1
   / \
  2   5
 / \   \
3   4   6</pre>

<p>将其展开为：</p>

<pre>1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6</pre>


 #### 题解
 **前序遍历方法**
 
 递归
 ```go
func flatten(root *TreeNode) {
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		prev,cur := list[i-1],list[i]
		prev.Left,prev.Right = nil,cur
	}
}

func preorder(root *TreeNode) []*TreeNode {
	var list []*TreeNode
	if root != nil {
		list = append(list, root)
		list = append(list, preorder(root.Left)...)
		list = append(list, preorder(root.Right)...)
	}
	return list
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 迭代
 ```go
func flatten(root *TreeNode) {
	var list []*TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			list = append(list, root)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	for i := 1; i < len(list); i++ {
		prev,cur := list[i-1],list[i]
		prev.Left,prev.Right = nil,cur
	}
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 **前驱和展开同时进行**
 使用方法一的前序遍历，由于将节点展开之后会破坏二叉树的结构而丢失子节点的信息，因此前序遍历和展开为单链表分成了两步。
 能不能在不丢失子节点的信息的情况下，将前序遍历和展开为单链表同时进行？
 
 之所以会在破坏二叉树的结构之后丢失子节点的信息，是因为在对左子树进行遍历时，没有存储右子节点的信息，
 在遍历完左子树之后才获得右子节点的信息。只要对前序遍历进行修改，在遍历左子树之前就获得左右子节点的信息，
 并存入栈内，子节点的信息就不会丢失，就可以将前序遍历和展开为单链表同时进行。
 ```go
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    stack := []*TreeNode{root}
    var prev *TreeNode
    for len(stack) > 0 {
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if prev != nil {
            prev.Left, prev.Right = nil, curr
        }
        left, right := curr.Left, curr.Right
        if right != nil {
            stack = append(stack, right)
        }
        if left != nil {
            stack = append(stack, left)
        }
        prev = curr
    }
}
```
 时间复杂度O(n),空间复杂度O(n)

 **寻找前驱节点**
 ```go
func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			prev := next
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = curr.Right
			curr.Left,curr.Right = nil,next
		}
		curr = curr.Right
	}
}
```
 时间复杂度O(n),空间复杂度O(1)