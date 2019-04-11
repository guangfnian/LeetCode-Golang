package p0020

func isLeft(x rune) bool {
	if x == '(' || x == '[' || x == '{' {
		return true
	}
	return false
}

type stack struct {
	s []rune
	l int
	c int
}

func (s *stack) Init() {
	s.s = make([]rune, 0)
	s.l = -1
	s.c = -1
}

func (s *stack) Push(x rune) {
	if s.l == s.c {
		s.s = append(s.s, x)
		s.c++
	} else {
		s.s[s.l+1] = x
	}
	s.l++
}

func (s *stack) Empty() bool {
	return s.l == -1
}

func (s *stack) Top() rune {
	if s.l == -1 {
		return '0'
	}
	return s.s[s.l]
}

func (s *stack) Pop() rune {
	if s.l == -1 {
		return '0'
	}
	s.l--
	return s.s[s.l+1]
}

func isPair(x, y rune) bool {
	if x == '(' && y == ')' {
		return true
	}
	if x == '[' && y == ']' {
		return true
	}
	if x == '{' && y == '}' {
		return true
	}
	return false
}

func isValid(s string) bool {
	var st stack
	st.Init()
	for _, i := range s {
		if isLeft(i) {
			st.Push(i)
		} else {
			if !isPair(st.Top(), i) {
				return false
			}
			st.Pop()
		}
	}
	return st.Empty()
}
