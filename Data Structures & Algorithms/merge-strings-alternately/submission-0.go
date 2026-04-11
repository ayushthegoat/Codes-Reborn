func mergeAlternately(word1 string, word2 string) string {
    i := 0
    j := 0
    res := ""
    for {
        if ((i < len(word1)) && (j < len(word2))) {
           res = res + string(word1[i]) 
           i++
           res = res + string(word2[j])
           j++
        } else {
            break
        }
    }
    for {
        if i < len(word1) {
            res = res + string(word1[i])
            i++
        } else {
            break
        }
    }
    for {
        if j < len(word2) {
            res = res + string(word2[j])
            j++
        } else {
            break
        }
    }
    return res
}
