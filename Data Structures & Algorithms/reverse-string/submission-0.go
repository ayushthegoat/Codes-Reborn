func reverseString(s []byte) {
    i := 0
    j := len(s) - 1
    for (i < j) {
        arb := s[i]
        s[i] = s[j]
        s[j] = arb
        i++
        j--
    }
}
