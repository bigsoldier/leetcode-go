package code

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var m = make(map[*Node]*Node)
	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}
	p = head
	for p != nil {
		node := m[p]
		if p.Next != nil {
			node.Next = m[p.Next]
		}
		if p.Random != nil {
			node.Random = m[p.Random]
		}
		p = p.Next
	}
	return m[head]
}
