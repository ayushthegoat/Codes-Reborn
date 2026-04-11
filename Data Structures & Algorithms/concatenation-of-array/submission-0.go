func getConcatenation(nums []int) []int {
    ans:= make([]int, 0, 2*len(nums)) 
        for _,val:= range nums {
           ans = append(ans, val)
        }
        for _,val:= range nums {
           ans = append(ans, val)
        }
    return ans
}
