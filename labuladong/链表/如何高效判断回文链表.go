package main

import "fmt"

// 寻找回文串：核心是从中心向两端扩展
func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// 判断是否是回文串
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

var left *listNode

// 判断链表中的数字是否是回文
func isPalindrome1(head *listNode) bool {
	left = head
	return traverse(head)
}

// 后续遍历：时间复杂度O(n),空间复杂度O(n)
func traverse(right *listNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.next)
	// 后续遍历
	res = res && (right.val == left.val)
	left = left.next
	return res
}

// 优化空间复杂度：快慢指针
func isPalindrome2(head *listNode) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	// 如果fast指针没有指向nil，说明链表长度为奇数，slow需要再前进一步
	if fast != nil {
		slow = slow.next
	}
	// 反转slow后面的链表
	left, right := head, reverse1(slow)
	for right != nil {
		if left.val != right.val {
			return false
		}
		left = left.next
		right = right.next
	}
	return true
}

// 优化空间复杂度：快慢指针(还原链表)
func isPalindrome3(head *listNode) bool {
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	// 反转slow后面的链表
	left, right := head, reverse1(slow.next)
	for right != nil {
		if left.val != right.val {
			return false
		}
		left = left.next
		right = right.next
	}
	// 还原链表
	slow.next = reverse1(right)
	return true
}

func main() {
	fmt.Println(palindrome("abba", 0, 3))
}
