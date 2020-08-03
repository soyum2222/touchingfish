package main

import (
	"fmt"
	"math"
)

//魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，
//在数组A中找出一个魔术索引，如果没有，则返回-1。若有多个魔术索引，返回索引值最小的一个。
//
//示例1:
//
// 输入：nums = [0, 2, 3, 4, 5]
// 输出：0
// 说明: 0下标的元素为0
//示例2:
//
// 输入：nums = [1, 1, 1]
// 输出：1
//提示:
//
//nums长度在[1, 1000000]之间
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/magic-index-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	fmt.Println(findMagicIndex([]int{-531369933, -469065528, -430059048, -428981853, -319235969, -288076332, -286667432, -282312559, -197049680, -197022263, -174416117, -138027773, -121899023, -111631966, -107567458, -70437707, -52463072, -45519851, -38641451, -15825815, -3835472, -1525043, 22, 566842886, 593757472, 605439236, 619794079, 640069993, 657657758, 718772950, 815849552, 839357142, 936585256, 1006188278, 1042347147, 1057129320, 1062178586, 1069769438}))
}
func findMagicIndex(nums []int) int {

	for i := 0; i < len(nums); {
		if nums[i] == i {
			return i
		} else {
			i = int(math.Max(float64(nums[i]), float64(i+1)))
		}
	}
	return -1
}
