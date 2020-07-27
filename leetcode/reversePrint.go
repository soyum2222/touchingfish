package main

import "fmt"

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
// 
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]
// 
//
//限制：
//
//0 <= 链表长度 <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(reversePrint(factory([]int{1, 2, 3})))
}

// for test
func factory(i []int) *ListNode {

	var head *ListNode
	var node *ListNode
	head = node
	for k, v := range i {

		if k == 0 {
			temp := &ListNode{
				Val:  v,
				Next: nil,
			}
			node = temp
			head = temp
		} else {
			node.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			node = node.Next
		}

	}

	return head
}

func reversePrint(head *ListNode) []int {
	var i []int
	head = reverse(head)

	for {
		if head != nil {
			i = append(i, head.Val)
			head = head.Next
		} else {
			return i
		}
	}
}

func reverse(head *ListNode) *ListNode {

	var last *ListNode

	for {

		if head != nil {

			var next *ListNode

			next = head.Next

			head.Next = last

			last = head
			if next == nil {
				return head
			}
			head = next

		} else {
			return nil
		}
	}
}
