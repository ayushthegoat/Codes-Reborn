/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0  
    }
    
    // Path through root
    throughRoot := maxDepth(root.Left) + maxDepth(root.Right)
    
    // Diameter in left subtree (might be larger)
    leftDiameter := diameterOfBinaryTree(root.Left)
    
    // Diameter in right subtree (might be larger)
    rightDiameter := diameterOfBinaryTree(root.Right)
    
    // Return the maximum of all three
    return max(throughRoot, max(leftDiameter, rightDiameter))
}
func maxDepth(root *TreeNode) int {
   if root == nil {
      return 0
   }
   return 1+ max(maxDepth(root.Left), maxDepth(root.Right))
}