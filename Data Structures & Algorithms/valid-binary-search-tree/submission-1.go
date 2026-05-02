/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {   
	 if root == nil {
		return true
	 }
	 return dfs(root, math.MinInt64, math.MaxInt64)
	
}
func dfs(root *TreeNode, mini int, maxi int) bool {
   if root == nil {
	return true
   }
   if root.Val <= mini || root.Val >= maxi {
   return false
   }
   return dfs(root.Left, mini, root.Val) && dfs(root.Right, root.Val, maxi)
}
