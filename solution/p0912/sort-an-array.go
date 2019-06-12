package p0912

func QuickSort(a []int, l, r int) {
	if l < r {
		pV := a[l]
		a[l], a[r] = a[r], a[l]
		cur := l
		for i := l; i < r; i++ {
			if a[i] < pV {
				a[cur], a[i] = a[i], a[cur]
				cur++
			}
		}
		a[cur], a[r] = a[r], a[cur]
		QuickSort(a, l, cur)
		QuickSort(a, cur+1, r)
	}
}

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}
