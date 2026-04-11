func removeElement(nums []int, val int) int {
    tracker:=0
    ans:=0
    for _,value := range nums {
        if value != val {
            nums[tracker] = value
            tracker++
            ans++
        } 
    }
    return ans

}
