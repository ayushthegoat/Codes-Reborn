func twoSum(nums []int, target int) []int {
    countMap := make(map[int]int) 
    for i, val:= range nums {
        compl := target - val
        if _,ok:=countMap[compl];ok {
            return []int{countMap[compl], i}
        } else {
             countMap[val] = i
        }
    }
    return []int{}
}
