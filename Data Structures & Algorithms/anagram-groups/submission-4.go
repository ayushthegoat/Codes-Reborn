func groupAnagrams(strs []string) [][]string {
     ans := [][]string{}
     mapper := make(map[string][]string)
     //here we are making the map of the key and then in the next line
     //i am checking for the existense of sorted version of the string if presetn in the map
     //if the sorted version is present it will append then current slice
     //so if act, cat so sorted will become act 
     // map[act] = append(act) ---> act --> [act]
     // map[act] = append(cat) ---> act --> [act, cat]
     // this way we keep filling by using the SORTED VERSION OF THE STRING as the key and then 
     //comaring it
     for i, val := range strs {
        sorted := sortString(strs[i])
        mapper[sorted] = append(mapper[sorted], val)
     }
    //  for i,val := range strs {
    //     sorted:= sortString(strs[i])
    //     if _, found := mapper[sorted]; found {
           
    //     }
    //  }

     for _, val := range mapper {
        ans = append(ans, val)
     }
     return ans
}
func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}
// func isAnagram(string s, string t) bool {
//     if len(s) != len(t) {
//         return false
//     }
//     var freq [26]int
//     for i, val := range s {
//         freq[s[i] - 'a']++
//         freq[t[i] - 'a']--

//         if freq[t[i] - 'a'] < 0 {
//             return false
//         }
//     }
//     return true
// }
