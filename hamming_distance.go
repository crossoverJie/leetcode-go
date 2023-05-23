package leetcode_go

import "math/bits"

// https://leetcode.cn/problems/hamming-distance/ 先异或，不同的位置会变为1，然后计算1的次数
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
