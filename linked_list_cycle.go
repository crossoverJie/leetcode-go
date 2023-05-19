package leetcode_go

// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	next := head.Next
	nnext := next.Next
	for nnext != nil {
		if next == nnext {
			return true
		}
		next = next.Next
		if nnext.Next == nil || nnext.Next.Next == nil {
			return false
		}
		nnext = nnext.Next.Next
	}

	return false
}
