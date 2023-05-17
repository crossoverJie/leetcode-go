package leetcode_go

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list14 := &ListNode{
		Val:  4,
		Next: nil,
	}
	list12 := &ListNode{
		Val:  2,
		Next: list14,
	}
	list11 := &ListNode{
		Val:  1,
		Next: list12,
	}

	list24 := &ListNode{
		Val:  4,
		Next: nil,
	}
	list23 := &ListNode{
		Val:  3,
		Next: list24,
	}
	list21 := &ListNode{
		Val:  1,
		Next: list23,
	}

	lists := mergeTwoLists(list11, list21)
	for lists.Next != nil {
		fmt.Printf("%d-> ", lists.Val)

		lists = lists.Next
	}
}
