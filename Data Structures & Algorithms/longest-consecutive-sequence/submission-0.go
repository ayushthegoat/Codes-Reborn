func longestConsecutive(nums []int) int {
   res:= 0
   counterMap := make(map[int]struct{})
   for _, num := range nums {
      counterMap[num] = struct{}{}
   }
   
   for _,num := range nums {
      streak := 0
      current := num
      for {
        if _, ok := counterMap[current]; !ok {
          break
        }
        streak++
        current++
      }

      if streak > res {
        res = streak
      }     
   }
   return res

}
