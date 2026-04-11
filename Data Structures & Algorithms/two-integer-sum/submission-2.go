func twoSum(nums []int, target int) []int {
    countMap := make(map[int]int) 
    for i, val:= range nums {

        //we are baisically checking if "iska compliment pehle se map mein hai kya ji????"
        //if hai then okay two sum hojayega if not store kro aage badho !
        //if target - currentNumber = someNumber && someNumber is persent in the MAP  
        //then someNumber + currentNumber = target 
        //meaning that is the two sum we are finding
        compl := target - val
        if j,ok:=countMap[compl];ok {
            return []int{j, i}
        } else {
             countMap[val] = i
        }
    }
    return []int{}
}
