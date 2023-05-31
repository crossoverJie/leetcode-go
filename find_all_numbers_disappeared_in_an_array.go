package leetcode_go

import "math"

// https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array
//func findDisappearedNumbers(nums []int) []int {
//    target := make(map[int]int, len(nums))
//    for i := 0; i < len(nums); i++ {
//        target[i+1] = i + 1
//    }
//    for _, num := range nums {
//        i, ok := target[num]
//        if ok {
//            delete(target, i)
//        }
//    }
//    r := make([]int, 0)
//    for _, v := range target {
//        r = append(r, v)
//    }
//    return r
//}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		i := int(math.Abs(float64(num)))
		if nums[i-1] > 0 {
			nums[i-1] = -nums[i-1]
		}
	}
	r := make([]int, 0)
	for i, num := range nums {
		if num > 0 {
			r = append(r, i+1)
		}
	}
	return r
}
