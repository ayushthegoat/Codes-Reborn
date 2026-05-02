/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	return height(root) != -1
	// lh := height(root.Left)
	// rh := height(root.Right)
	// if abs(lh - rh) > 1 {
	// 	return false
	// }
	// lb := isBalanced(root.Left)
	// rb := isBalanced(root.Right)
	// return lb && rb
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	//optimization we directly return the balance factor 
	if lh == -1 {
		return -1
	}
	rh := height(root.Right)
	//optimization we directly return the balance factor 
	if rh == -1 {
		return -1
	}
	if abs(lh - rh) > 1 {
		return -1
	}
	return 1 + max(lh, rh)
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}