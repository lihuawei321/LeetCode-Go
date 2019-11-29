package main

//给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := inorder(root)
	return arr[k-1]
}

//方法1
func inorder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	var arr []int
	left := inorder(node.Left)
	right := inorder(node.Right)
	arr = append(arr, left...)
	arr = append(arr, node.Val)
	arr = append(arr, right...)

	return arr

}

//方法2
func kthSmallest2(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder2(root, k, &count, &res)
	return res
}

func inorder2(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder2(node.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = node.Val
			return
		}
		inorder2(node.Right, k, count, ans)
	}
}
func main() {

}
