type Pair struct {
  key int
  val int
}
func topKFrequent(nums []int, k int) []int {
     mapper := make(map[int]int) 
     for i := range nums {
       mapper[nums[i]]++
     }
     pairs := make([]Pair, len(mapper))
     for key, value := range mapper {
        pairs = append(pairs, Pair{key: key, val: value})
     }
     sort.Slice(pairs, func(i int, j int) bool {
        return pairs[i].val > pairs[j].val
     })
     ans := make([]int,0, k)
     for i:=0; i<k; i++ {
      ans = append(ans, pairs[i].key)
     }
     return ans
}
