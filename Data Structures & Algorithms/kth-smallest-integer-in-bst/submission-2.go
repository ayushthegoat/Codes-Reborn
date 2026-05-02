/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    s := []*TreeNode{}
	curr := root
   //iterative dfs approach we iterate using a manul stack 
	for len(s) > 0 || curr != nil {
		for curr != nil {
			s = append(s , curr)
			curr = curr.Left
		}
		// we first assing the currents stacks last meaning the leftmost element 
		//we do this because the curr will become nil after the last recursive left for loop
		curr = s[len(s) - 1]
		// then we pop the last from the stack cause that we are processing 
		s = s[:len(s) - 1]

		k--
		//decreatse teh val of k when it matches we have reached the kth element 
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
	return 0
}
// dfs appraoch we do normal dfs but keep track of the times var whenerv it reaches the val we return 
// func dfs(root *TreeNode, k int, times *int) int {
// 	if root == nil {
// 		return -1
// 	}
//     leftVal := dfs(root.Left, k, times)
// 	if leftVal != -1 {
// 		return leftVal
// 	}
//     *times++
// 	if k == *times {
// 		return root.Val
// 	}
// 	return dfs(root.Right, k, times)
// }
