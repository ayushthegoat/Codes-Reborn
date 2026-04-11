// func longestCommonPrefix(strs []string) string {
//     if len(strs) == 0 {
//         return ""
//     }
//     ans := ""
//     s := findSmallestString(strs)
//     for i := range s {               
//         for _, jval := range strs {
//             if s[i] != jval[i] {
//                 return ans
//             }
//         }
//         ans = ans + string(s[i])   
//     }

//     return ans                     
// }
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }

    // We can use strs[0] as reference (or the shortest string)
    for i := 0; i < len(strs[0]); i++ {
        ch := strs[0][i]

        // check if this character appears at position i in ALL other strings
        for j := 1; j < len(strs); j++ {
            // if string is shorter or char doesn't match → stop
            if i >= len(strs[j]) || strs[j][i] != ch {
                return strs[0][:i]
            }
        }
    }

    // all characters matched up to the end of strs[0]
    return strs[0]
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
