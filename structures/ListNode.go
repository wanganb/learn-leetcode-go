package structures

import (
	"strconv"
	"strings"
)

//单链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

func Ints2ListNode(ints []int) *ListNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &ListNode{Val: ints[0]}

	node := root

	for i := 1; i < n; i++ {
		node.Next = &ListNode{Val: ints[i]}
		node = node.Next
	}

	return root
}

func (l *ListNode) String() string {

	s := "["
	for {
		s = s + strconv.Itoa(l.Val) + ","
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return strings.Trim(s, ",") + "]"
}
