type MyHashSet struct {
    store []int
}

func Constructor() MyHashSet {
    return MyHashSet {
      store: make([]int, 1000000),
    
    }
}

func (this *MyHashSet) Add(key int) {
    if this.store[key] > 0 {
      return 
    } else {
      this.store[key]++
    }
    
}

func (this *MyHashSet) Remove(key int) {
    if this.store[key] > 0 {
        this.store[key]--
    }
}

func (this *MyHashSet) Contains(key int) bool {
    return this.store[key] > 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 