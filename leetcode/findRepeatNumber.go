package main

import "fmt"

//找出数组中重复的数字。
//
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
//示例 1：
//
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
// 
//
//限制：
//
//2 <= n <= 100000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 这题我在书上看过，思路为把值移动到他作为索引的位子，如果该位子有和他相等的值，那么就重复 这样最大O（n）复杂度
func main() {
	fmt.Println(findRepeatNumber([]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
}
func findRepeatNumber(nums []int) int {

	i := 0
loop:
	index := nums[i]

	if i == index {
		i++
		goto loop
	}

	if index == nums[index] {
		return index
	} else {
		index, nums[index] = nums[index], index
	}
	goto loop

}
