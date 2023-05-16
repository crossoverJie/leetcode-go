package leetcode_go

func twoSum(nums []int, target int) []int {
	r := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for l := i + 1; l < len(nums); l++ {
			b := nums[l]
			if a+b == target {
				r[0], r[1] = i, l
				return r
			}
		}
	}
	return r
}
