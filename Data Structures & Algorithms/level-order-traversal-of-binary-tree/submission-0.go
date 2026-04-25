/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
   if root == nil {
      return [][]int{}
   }
      res := [][]int{}
      q := []*TreeNode{root}
      
     
      for len(q) > 0 {
         current := []int{}
         w := len(q)
         
         for i:=0; i<w; i++ {
           node := q[0]
           q = q[1:]
           current = append(current, node.Val)

           if node.Left != nil {
            q = append(q, node.Left)
           }

            if node.Right != nil {
            q = append(q, node.Right)
           }
         }
         res = append(res, current)
        
      }

      return res

}
