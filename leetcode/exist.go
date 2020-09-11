package main

import "fmt"

//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
//
//[["a","b","c","e"],
//["s","f","c","s"],
//["a","d","e","e"]]
//
//但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
//
//
//
//示例 1：
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//示例 2：
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//提示：
//
//1 <= board.length <= 200
//1 <= board[i].length <= 200
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var maxy, maxx int

func exist(board [][]byte, word string) bool {

	if word == "" {
		return false
	}

	maxx = len(board)
	if maxx == 0 {
		return false
	}

	maxy = len(board[0])

	for x, v := range board {

		for y, vv := range v {

			if vv == word[0] {

				if len(word) == 1 {
					return true
				}

				if move(x, y, word, 1, board) {
					return true
				}

			}

		}
	}

	return false

}

func move(x, y int, word string, index int, board [][]byte) bool {
	tmp := board[x][y]
	board[x][y] = ' '
	for i := 0; i < 4; i++ {

		fx, fy := x, y
		switch i % 4 {
		case 0:
			fx, fy = up(x, y)
		case 1:
			fx, fy = right(x, y)
		case 2:
			fx, fy = down(x, y)
		case 3:
			fx, fy = left(x, y)

		}

		if fx >= maxx || fy >= maxy || fx < 0 || fy < 0 {
			continue
		}

		if word[index] == board[fx][fy] {
			index++
			if index >= len(word) {
				return true
			}
			if move(fx, fy, word, index, board) {
				return true
			} else {
				index--
			}
		}
	}

	board[x][y] = tmp
	return false
}

func up(x, y int) (int, int) {

	return x + 1, y
}

func left(x, y int) (int, int) {

	return x, y - 1
}

func right(x, y int) (int, int) {

	return x, y + 1
}

func down(x, y int) (int, int) {
	return x - 1, y
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))        //true
	fmt.Println(exist([][]byte{{'A', 'B'}, {'C', 'D'}}, "ABCD"))                                                    //false
	fmt.Println(exist([][]byte{{'A', 'B'}}, "BA"))                                                                  //true
	fmt.Println(exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB"))                          //true
	fmt.Println(exist([][]byte{{'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A'}}, "AAAAAAAAAAAAA")) //false
}
