package p0012

func intToRoman(num int) string {
	ret := ""
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	cur := 0
	for num > 0 {
		for num >= numbers[cur] {
			num -= numbers[cur]
			ret += strs[cur]
		}
		cur++
	}
	return ret
}
