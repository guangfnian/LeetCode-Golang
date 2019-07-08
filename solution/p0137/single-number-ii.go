package p0137

func singleNumber(nums []int) int {
	var ret int
	var count, bit int
	for i := 31; i >= 0; i-- {
		count = 0
		bit = 1 << uint(i)
		for _, n := range nums {
			if n&bit != 0 {
				count++
			}
		}
		if count%3 != 0 {
			ret |= bit
		}
	}
	return int(int32(ret))
}
