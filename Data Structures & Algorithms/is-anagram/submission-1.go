func isAnagram(s string, t string) bool {
    counter := make(map[rune]int)
    if len(s) != len(t) {
        return false
    }
     for _, val := range s {
        counter[val]++
     }
     for _, val := range t {
        if _, ok := counter[val]; ok {
            counter[val]--
            if counter[val] < 0 {
                return false
            }
        } else {
            return false
        }
     }
     return true
}
