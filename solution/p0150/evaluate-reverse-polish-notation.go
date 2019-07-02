package p0150

import "strconv"

type Stack struct {
	s []int
}

func (s *Stack) Init() {
	s.s = make([]int, 0)
}

func (s *Stack) Empty() bool {
	return len(s.s) == 0
}

func (s *Stack) Push(x int) {
	s.s = append(s.s, x)
}

func (s *Stack) Pop() int {
	// check empty
	ret := s.s[len(s.s)-1]
	s.s = s.s[0 : len(s.s)-1]
	return ret
}

func evalRPN(tokens []string) int {
	s := Stack{}
	s.Init()
	var x, y, i int
	for _, t := range tokens {
		switch t {
		case "+":
			x = s.Pop()
			y = s.Pop()
			s.Push(y + x)
		case "-":
			x = s.Pop()
			y = s.Pop()
			s.Push(y - x)
		case "*":
			x = s.Pop()
			y = s.Pop()
			s.Push(y * x)
		case "/":
			x = s.Pop()
			y = s.Pop()
			s.Push(y / x)
		default:
			i, _ = strconv.Atoi(t)
			s.Push(i)
		}
	}
	return s.Pop()
}
