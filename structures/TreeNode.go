package structures

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	fmt.Println(queue)
	queue[0] = root
	fmt.Println(queue)

	i := 1

	for i < n {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(queue)

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{
				Val: ints[i],
			}
			queue = append(queue, node.Left)
			fmt.Println(queue)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{
				Val: ints[i],
			}
			queue = append(queue, node.Right)
			fmt.Println(queue)
		}
		i++
	}

	return root
}
