func hasDuplicate(nums []int) bool {
    mapper:= make(map[int]struct{});
    for _,val := range nums {
        if _, ok :=  mapper[val]; ok {
            return true
        }
        mapper[val] = struct{}{}
    }
    return false
}
