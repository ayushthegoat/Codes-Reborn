type MyHashMap struct {
    store[] int
}

func Constructor() MyHashMap {
   
    s :=  make([]int, 1_000_001)
    for i := range s {
      s[i] = -1
   }
    return MyHashMap{
        store : s,
    } 
}

func (this *MyHashMap) Put(key int, value int) {
        this.store[key] = value
}

func (this *MyHashMap) Get(key int) int {
    if this.store[key] != -1 {
        return this.store[key] 
    } else {
        return -1
    }
}

func (this *MyHashMap) Remove(key int) {
        this.store[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */