package p0001

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, p := range nums {
		if j, ok := mp[target-p]; ok {
			return []int{j, i}
		}
		mp[p] = i
	}
	return nil
}