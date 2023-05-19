package leetcode_go

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {

	t15 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	t7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	t20 := &TreeNode{
		Val:   20,
		Left:  t15,
		Right: t7,
	}
	t9 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	t3 := &TreeNode{
		Val:   3,
		Left:  t9,
		Right: t20,
	}
	fmt.Println(maxDepth(t3))

}
