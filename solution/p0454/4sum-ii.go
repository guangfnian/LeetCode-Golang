package p0454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	mpa := make(map[int]int)
	for i := range A {
		for j := range B {
			mpa[A[i]+B[j]]++
		}
	}
	ret := 0
	for i := range C {
		for j := range D {
			if k, ok := mpa[-C[i]-D[j]]; ok {
				ret += k
			}
		}
	}
	return ret
}
