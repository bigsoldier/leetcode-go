package code

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var visited = make(map[*Node]*Node)
	var visit func(*Node) *Node
	visit = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if v, ok := visited[node]; ok {
			return v
		}
		cloneNode := &Node{node.Val, []*Node{}}
		visited[node] = cloneNode
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, visit(n))
		}
		return cloneNode
	}
	return visit(node)
}
