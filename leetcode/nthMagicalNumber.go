package main

import (
	"fmt"
	"math"
)

//如果正整数可以被 A 或 B 整除，那么它是神奇的。
//
//返回第 N 个神奇数字。由于答案可能非常大，返回它模 10^9 + 7 的结果。
//
// 
//
//示例 1：
//
//输入：N = 1, A = 2, B = 3
//输出：2
//示例 2：
//
//输入：N = 4, A = 2, B = 3
//输出：6
//示例 3：
//
//输入：N = 5, A = 2, B = 4
//输出：10
//示例 4：
//
//输入：N = 3, A = 6, B = 4
//输出：8
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/nth-magical-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var mod = int(math.Pow(10, 9) + 7)

func main() {
	fmt.Println(nthMagicalNumber(1, 2, 3)) //2
	fmt.Println(nthMagicalNumber(3, 6, 4)) //8
	fmt.Println(nthMagicalNumber(3, 3, 8)) //8
}

func nthMagicalNumber(N int, A int, B int) int {

	var lcm int
	var spacing int

	var foo int
	var i int = 0
	for ; foo%A != 0 || foo%B != 0 || foo == 0; i++ {
		foo = A * (i + 1)
	}
	lcm = foo

	foo = 0
	var j int = 0
	for ; foo%A != 0 || foo%B != 0 || foo == 0; j++ {
		foo = B * (j + 1)
	}

	spacing = i + j - 1

	if N%spacing == 0 {
		return (lcm * N / spacing) % mod
	} else {

		n := (N / spacing) * spacing

		magic := lcm * n / spacing
		nmagic := magic
		var countA = 1
		var countB = 1
		for ; n < N; {

			if magic+(A*countA) > magic+(B*countB) {
				nmagic = magic+(B*countB)
				countB++
			} else {
				nmagic = magic+(A*countA)
				countA++
			}

			n++
		}

		return nmagic % mod

	}

}
