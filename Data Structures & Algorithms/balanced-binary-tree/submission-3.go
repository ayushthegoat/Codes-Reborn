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
    // BASE CASE: Agar node nil hai, toh uski height 0 hoti hai
    // Aur nil node always balanced hota hai
    if root == nil {
        return 0
    }
    
    // STEP 1: Pehle left subtree ki height nikaalo
    // Ye recursively left side mein deepest tak jaayega
    lh := height(root.Left)
    
    // OPTIMIZATION 1: Agar left subtree already unbalanced hai
    // Matlab left se -1 aa gaya (jo ki imbalance ka signal hai)
    // Toh further check karne ki zaroorat nahi, seedha -1 return karo
    // Issey hum early exit kar lete hain, time bachta hai
    if lh == -1 {
        return -1  // Left subtree hi unbalanced hai, poora tree unbalanced
    }
    
    // STEP 2: Ab right subtree ki height nikaalo
    // Tabhi right mein jaayenge jab left balanced tha
    rh := height(root.Right)
    
    // OPTIMIZATION 2: Agar right subtree unbalanced hai
    // Toh seedha -1 return, no need to check further
    if rh == -1 {
        return -1  // Right subtree unbalanced, poora tree unbalanced
    }
    
    // STEP 3: Current node pe balance check karo
    // Formula: |left height - right height| <= 1 honi chahiye
    // Agar difference 1 se zyada hai, toh current node pe imbalance hai
    if abs(lh - rh) > 1 {
        return -1  // Is node pe unbalanced ho gaya
    }
    
    // STEP 4: Sab kuch balanced hai toh height return karo
    // Height = 1 (current node) + max(left subtree height, right subtree height)
    // Ye parent ko batayega ki "main kitna deep hoon"
    return 1 + max(lh, rh)

    // -1 ka matlab: "Bhaiyaji yaha pe khela kahtam hai , yahan se neeche unbalanced hai, aage mat dekh"
    // Positive number ka matlab: "Sab sahi hai, meri height itni hai"
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}