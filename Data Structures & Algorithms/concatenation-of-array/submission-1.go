func getConcatenation(nums []int) []int {
    ans:= make([]int,  2*len(nums)) 
        for i ,val:= range nums {
           ans[i]  = val
           ans[i + len(nums)] = val
        }
    
    return ans
}
