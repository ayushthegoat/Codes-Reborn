func isPalindrome(s string) bool {
    sanitize := sanitize(s)
    //firse we sanitize so to remove unwanted characted like special characters 
    i := 0
    j := len(sanitize)-1
    for {
        if i < j {
            if !strings.EqualFold(string(sanitize[i]), string(sanitize[j])) {
                return false
            }
            i++
            j--
        } else {
            break
        }
    }
    return true
}
//this fucntion removed the special charactes from ower string
func sanitize(str string) string {
    ans:=""
    for _, ch := range str {
        if ((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')) {
            ans = ans + string(ch)
        }
    }
    return ans
}
