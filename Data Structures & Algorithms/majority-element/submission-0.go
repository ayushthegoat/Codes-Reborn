func majorityElement(nums []int) int {
    counter := make(map[int]int) 
    for _, val := range nums {
        counter[val]++
        if counter[val] > len(nums)/2 {
            return val
        }
    }
    return 0
}
