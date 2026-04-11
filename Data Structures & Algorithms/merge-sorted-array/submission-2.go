// this appraoch is also correct but requires an extra array to store the elements 
// func merge(nums1 []int, m int, nums2 []int, n int) {
//    i := 0
//    j := 0
//    t := 0
//    ans := make([]int, m + n)
//    for (i < m) && (j < n) {
//        if nums1[i] <= nums2[j] {
//            ans[t] = nums1[i]
//            i++
//            t++
//        } else {
//            ans[t] = nums2[j]
//            j++
//            t++
//        }
//    }
//    for (i < m) {
//     ans[t] = nums1[i]
//     i++
//     t++
//    }
//    for (j < n) {
//      ans[t] = nums2[j]
//      j++
//      t++
//    }
//    for i := range nums1 {
//      nums1[i] = ans[i]
//    }
//    //return ans
// }
// this function is used to sort it wihtout using any extra space 
//because the first array already contains the extra 0s in the last 
//so we take 3 pointers but in reverse
// one storage pointer marks the last of the array
//so for example
//[10,20,20,40,0,0]  [1,2]
//          m    k      n 
//in first loop
// if m[i]  > n[j]
//  k = m[i] so k = 60
// second loop
//[10,20,20,40,40,60]
//  k = m[i] so k = 40
//[10,20,20,20,40,60]
//[10,20,10,20,40,60]
// then loop ends 
// then the other loops kick in 
//    for (j >= 0) {
//      nums1[k] = nums2[j]
//      j--
//      k--
//    }
// this fills the remaining from nums[2] to tha first of the array
////[10,2,20,20,40,60]
//[1,2,10,20,40,60]
//above array is the answer
func merge(nums1 []int, m int, nums2 []int, n int) {
   i := m-1
   j := n-1
   k := m + n - 1
   for i >= 0 && j >= 0  {
      if nums1[i] > nums2[j] {
        nums1[k] = nums1[i]
        i--
        k--
      } else {
        nums1[k] = nums2[j]
        j--
        k--
      }
   }
  for (i >= 0) {
    nums1[k] = nums1[i]
    i--
    k--
   }
   for (j >= 0) {
     nums1[k] = nums2[j]
     j--
     k--
   }



}