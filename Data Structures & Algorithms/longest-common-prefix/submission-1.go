func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    ans := ""
    s := findSmallestString(strs)
    for i := range s {               
        for _, jval := range strs {
            if s[i] != jval[i] {
                return ans
            }
        }
        ans = ans + string(s[i])   
    }

    return ans                     
}

func findSmallestString(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    smallest := strs[0] 

    for i := 1; i < len(strs); i++ {
        if strs[i] < smallest {
            smallest = strs[i]
        }
    }

    return smallest
}
// below is function to find common Prefix for 2 strings for ONLY TWO STRINGS
// func findPrefix(s string, t string) string {
//     cutOff := ""
//     i:= 0
//     j:= 0
//     for {
//      if (i < len(s) && j < len(t)) {
//         if s[i] == t[j] {
//            cutOff = cutOff + string(s[i])
//            i++
//            j++
//         } else {
//             break
//         }
//       } else {
//         break
//       }
//     }
//     return cutOff
// }
