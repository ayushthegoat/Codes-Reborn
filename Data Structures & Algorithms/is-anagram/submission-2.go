func isAnagram(s string, t string) bool {
    // counter := make(map[rune]int)
     if len(s) != len(t) {
         return false
     }
    // for _, val := range s {
    //     counter[val]++
    // }
    // for _, val := range t {
    //     if _, ok := counter[val]; ok {
    //         counter[val]--
    //         if counter[val] < 0 {
    //             return false
    //         }
    //     } else {
    //         return false
    //     }
    //  }
    //  return true

    //“I am using 'a' as the starting point in ASCII, so subtracting 
    //'a' from any lowercase character gives me a zero-based index.

    // this approach only works if incoming string is all lowercase if
    //its also contaisn uppercase then 
    //var freq [52]int
    //  if c >= 'a' && c <= 'z' {
    //     freq[c-'a']++
    // } else if c >= 'A' && c <= 'Z' {
    //     freq[c-'A'+26]++
    // }
    //then we use something like this where whe check for case senticiyt and then update 
    var freq[26]int
    for i := 0; i < len(s); i++ {
        freq[s[i] - 'a']++
        freq[t[i] - 'a']--
    }
    for _, val := range freq {
       if val != 0 {
          return false
       }
    }
    
     return true
}
