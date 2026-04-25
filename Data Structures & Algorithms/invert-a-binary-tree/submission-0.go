/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
   return inverter(root)
}
func inverter(root *TreeNode) *TreeNode {
   if root == nil {
      return nil
   }
   hold := root.Left
   root.Left = root.Right
   root.Right = hold
   inverter(root.Left)
   inverter(root.Right)
   return root
}
