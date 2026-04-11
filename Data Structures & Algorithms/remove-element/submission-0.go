func removeElement(nums []int, val int) int {
    for i, value := range nums {
        if value == val {
            nums[i] = -1
        }
    }
    tracker:=0
    ans:=0
    for _,val := range nums {
        if val != -1 {
            nums[tracker] = val
            tracker++
            ans++
        } 
    }
    return ans

}
