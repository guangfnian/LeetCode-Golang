package p0380

import "math/rand"

type RandomizedSet struct {
	arr []int
	mp  map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: make([]int, 0),
		mp:  make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.mp[val] = len(this.arr) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.mp[val]; ok {
		r := len(this.arr) - 1
		this.arr[v], this.arr[r] = this.arr[r], this.arr[v]
		this.mp[this.arr[v]] = v
		delete(this.mp, val)
		this.arr = this.arr[:r]
		return true
	} else {
		return false
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	r := rand.Int31n(int32(len(this.arr)))
	return this.arr[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
