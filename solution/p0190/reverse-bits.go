package p0190

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	one := uint32(1)
	for i := 0; i < 32; i++ {
		ret = (ret << one) | (num & one)
		num = num >> one
	}
	return ret
}
