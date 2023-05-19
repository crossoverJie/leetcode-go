package leetcode_go

import (
	"fmt"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	n4 := &ListNode{
		Val: 4,
	}
	n0 := &ListNode{
		Val:  0,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n0,
	}
	n4.Next = n2

	n3 := &ListNode{
		Val:  3,
		Next: n2,
	}
	fmt.Println(hasCycle(n3))
}
