package p0384

import "math/rand"

type Solution struct {
	ori []int
	cp  []int
}

func Constructor(nums []int) Solution {
	s := Solution{
		ori: nums,
		cp:  make([]int, len(nums)),
	}
	copy(s.cp, nums)
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.ori
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.cp)
	for i := int32(length - 1); i >= 0; i-- {
		idx := rand.Int31n(i + 1)
		this.cp[i], this.cp[idx] = this.cp[idx], this.cp[i]
	}
	return this.cp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
