package leetcode_go

import "math/bits"

// https://leetcode.cn/problems/counting-bits/?envType=featured-list&envId=2cktkvj
// 循环遍历计算每位二进制 1 的数量
func countBits(n int) []int {
	array := make([]int, n+1)
	for i := 0; i <= n; i++ {
		onesCount := bits.OnesCount(uint(i))
		array[i] = onesCount
	}
	return array
}
