package main

import (
	"fmt"
)

//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000007，请返回 1。
//
//示例 1：
//
//输入：n = 2
//输出：2
//示例 2：
//
//输入：n = 7
//输出：21
//示例 3：
//
//输入：n = 0
//输出：1
//提示：
//
//0 <= n <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(numWays(100))
}

// func2 fib
func numWays(n int) int {

	var a, b = 1, 1

	for i := 2; i <= n; i++ {
		b, a = a, (a+b)%1000000007
	}

	return a % 1000000007
}

// func1 formula
/**
func numWays(n int) int {

	towNum := n / 2

	total := towNum + n%2

	var r int64
	for towNum >= 0 {

		r += C(total, towNum)
		total++
		towNum--
	}

	return int(r % 1000000007)

}

func C(m, n int) int64 {

	return new(big.Int).Mod(new(big.Int).Div(factorial(m), new(big.Int).Mul(factorial(n), factorial(m-n))), big.NewInt(1000000007)).Int64()
}

func factorial(n int) *big.Int {
	m := big.NewInt(1)
	for i := 1; i <= n; i++ {
		m = m.Mul(m, big.NewInt(int64(i)))
	}

	return m
}
*/
