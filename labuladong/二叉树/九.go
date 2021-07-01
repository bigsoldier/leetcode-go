package main

import (
	"strconv"
	"strings"
)

// 二叉树的序列化和反序列化

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var s string
	serialize(root, &s)
	return s
}

func serialize(root *TreeNode, s *string) {
	if root == nil {
		*s += "#,"
		return
	}
	// 前序遍历
	*s += strconv.Itoa(root.val) + ","
	serialize(root.left, s)
	serialize(root.right, s)
}

// Deserializes your encoded data to tree.
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	return deserialize(&l)
}

func deserialize(data *[]string) *TreeNode {
	num := (*data)[0]
	*data = (*data)[1:]
	if num == "#" {
		return nil
	}
	val, _ := strconv.Atoi(num)
	root := &TreeNode{val: val}
	root.left = deserialize(data)
	root.right = deserialize(data)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
