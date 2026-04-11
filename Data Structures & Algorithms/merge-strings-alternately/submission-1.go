// func mergeAlternately(word1 string, word2 string) string {
//     i := 0
//     j := 0
//     res := ""
//     for {
//         if ((i < len(word1)) && (j < len(word2))) {
//            res = res + string(word1[i]) 
//            i++
//            res = res + string(word2[j])
//            j++
//         } else {
//             break
//         }
//     }
//     for {
//         if i < len(word1) {
//             res = res + string(word1[i])
//             i++
//         } else {
//             break
//         }
//     }
//     for {
//         if j < len(word2) {
//             res = res + string(word2[j])
//             j++
//         } else {
//             break
//         }
//     }
//     return res
// }

func mergeAlternately(word1 string, word2 string) string {
    var sb strings.Builder
  

    i, j := 0, 0
    for i < len(word1) && j < len(word2) {
        sb.WriteByte(word1[i])
        i++
        sb.WriteByte(word2[j])
        j++
    }

    // remaining part of word1
    for i < len(word1) {
        sb.WriteByte(word1[i])
        i++
    }

    // remaining part of word2
    for j < len(word2) {
        sb.WriteByte(word2[j])
        j++
    }

    return sb.String()
}
