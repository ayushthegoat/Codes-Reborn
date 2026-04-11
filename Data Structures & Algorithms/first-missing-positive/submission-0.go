func firstMissingPositive(nums []int) int {
     sort.Ints(nums)
     smallestPos := 1
     for _, num := range nums {
       if num < smallestPos {
        continue
       }

        if num == smallestPos {
          smallestPos++
        } else if num > smallestPos {
          return smallestPos
        }
     }
     return smallestPos
}
