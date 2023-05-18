package leetcode_go

var result map[int]int

// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if result == nil {
		result = make(map[int]int)
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	v, ok := result[n]
	if ok {
		return v
	} else {
		r := climbStairs(n-1) + climbStairs(n-2)
		result[n] = r
		return r
	}

}
