/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if root == nil {
		return root
	}
	valOfCur := root.Val
	if p.Val > valOfCur && q.Val > valOfCur {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < valOfCur && q.Val < valOfCur {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
	
}
