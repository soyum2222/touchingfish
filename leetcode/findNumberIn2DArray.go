package main

import "fmt"

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//
// 
//
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。
//
// 
//
//限制：
//
//0 <= n <= 1000
//
//0 <= m <= 1000
//
// 
//
//注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思路： 先从坐标（0，0）开始往右找，如果（x+1,y）> target ,那么就向下找
// 如果（x,y+1）> target 那么就往左找，即下一个坐标为（x-1,y+1）
func main() {

	foo := [][]int{{-5}}
	//foo := [][] int{{-1, 3}}

	fmt.Println(findNumberIn2DArray(foo, -10))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var x, y int

	for {

		i := matrix[x][y]

		if i == target {
			return true
		}

		if i < target {

			x, y = right(x, y)

			// roll back
			if x >= len(matrix) || matrix[x][y] > target {
				x, y = left(x, y)
				x, y = down(x, y)
				if y >= len(matrix[0]) {
					return false
				}
			}

		}

		if i > target {
			x, y = left(x, y)
			if y >= len(matrix[0]) {
				return false
			}
			if x < 0 {
				return false
			}

		}

	}

}

func down(x, y int) (int, int) {
	y++
	return x, y
}

func up(x, y int) (int, int) {
	y--
	return x, y
}

func right(x, y int) (int, int) {
	x++
	return x, y
}

func left(x, y int) (int, int) {
	x--
	return x, y
}
