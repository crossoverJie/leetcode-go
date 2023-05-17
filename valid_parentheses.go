package leetcode_go

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	dict := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if i > 0 && len(stack) > 0 {
			top := stack[len(stack)-1]
			if b == dict[top] {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, b)
			}
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
