func merge(nums1 []int, m int, nums2 []int, n int) {
   i := 0
   j := 0
   t := 0
   ans := make([]int, m + n)
   for (i < m) && (j < n) {
       if nums1[i] <= nums2[j] {
           ans[t] = nums1[i]
           i++
           t++
       } else {
           ans[t] = nums2[j]
           j++
           t++
       }
   }
   for (i < m) {
    ans[t] = nums1[i]
    i++
    t++
   }
   for (j < n) {
     ans[t] = nums2[j]
     j++
     t++
   }
   for i := range nums1 {
     nums1[i] = ans[i]
   }
   //return ans
}
