package p0008

func isDigit(x byte) bool {
	if x <= '9' && x >= '0' {
		return true
	}
	return false
}

func toInt64(x byte) int64 {
	return int64(x - '0')
}

func myAtoi(str string) int {
	const Max = 1<<31 - 1
	const Min = -1 << 31
	length := len(str)
	if length == 0 {
		return 0
	}
	cur := 0
	for cur < length && str[cur] == ' ' {
		cur++
	}
	if cur == length {
		return 0
	}
	if str[cur] != '-' && str[cur] != '+' && !isDigit(str[cur]) {
		return 0
	}
	flag := false
	if str[cur] == '-' {
		flag = true
	}
	if !isDigit(str[cur]) && (cur == length-1 || !isDigit(str[cur+1])) {
		return 0
	}
	var ret int64 = 0
	if isDigit(str[cur]) {
		ret = toInt64(str[cur])
	}
	cur++
	for ; cur < length && isDigit(str[cur]) && ret < Max; cur++ {
		ret = ret*10 + toInt64(str[cur])
	}
	if flag {
		ret = -ret
	}
	if ret > Max {
		return Max
	}
	if ret < Min {
		return Min
	}
	return int(ret)
}
