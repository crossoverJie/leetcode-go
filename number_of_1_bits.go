package leetcode_go

import "math/bits"

// https://leetcode.cn/problems/number-of-1-bits/ 计算二进制中 1 的次数，内置函数 OnesCount 实现
func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}
