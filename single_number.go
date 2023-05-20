package leetcode_go

// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	c := make(map[int]int)
	for _, num := range nums {
		v, ok := c[num]
		if ok {
			delete(c, v)
		} else {
			c[num] = num
		}
	}
	for _, v := range c {
		return v
	}
	return 0
}
