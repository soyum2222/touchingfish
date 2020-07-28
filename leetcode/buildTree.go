package main

//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
// 
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//        3
//       / \
//      9  20
//        /  \
//       15   7
// 
//
//限制：
//
//0 <= 节点个数 <= 5000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}


func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	tree := TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			tree.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
			tree.Right = buildTree(preorder[len(inorder[:i])+1:len(inorder[:i])+len(inorder[i+1:])+1], inorder[i+1:])
			return &tree
		}
	}

	return nil
}
