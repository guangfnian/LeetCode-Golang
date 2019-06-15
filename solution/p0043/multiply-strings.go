package p0043

func reverse(arr []rune) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func subChar(arr []rune) {
	for i := range arr {
		arr[i] -= '0'
	}
}

func addChar(arr []rune) {
	for i := range arr {
		arr[i] += '0'
	}
}

func multiply(num1 string, num2 string) string {
	a := []rune(num1)
	b := []rune(num2)
	reverse(a)
	reverse(b)
	subChar(a)
	subChar(b)
	la, lb := len(a), len(b)
	c := make([]rune, la+lb)
	for i := 0; i < la; i++ {
		for j := 0; j < lb; j++ {
			c[i+j] += a[i] * b[j]
		}
	}
	add := rune(0)
	for i := range c {
		c[i] += add
		if c[i] > 9 {
			add = c[i] / 10
			c[i] = c[i] % 10
		} else {
			add = 0
		}
	}
	r := len(c) - 1
	for r > 0 && c[r] == 0 {
		r--
	}
	c = c[:r+1]
	reverse(c)
	addChar(c)
	return string(c)
}
