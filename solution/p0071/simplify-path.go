package p0071

import "strings"

type Stack struct {
	s []string
}

func (s *Stack) Init() {
	s.s = make([]string, 0)
}

func (s *Stack) Len() int {
	return len(s.s)
}

func (s *Stack) Empty() bool {
	return len(s.s) == 0
}

func (s *Stack) Push(str string) {
	s.s = append(s.s, str)
}

func (s *Stack) Pop() string {
	if len(s.s) > 0 {
		ret := s.s[len(s.s)-1]
		s.s = s.s[0 : len(s.s)-1]
		return ret
	}
	return ""
}

func (s *Stack) LPop() string {
	if len(s.s) > 0 {
		ret := s.s[0]
		s.s = s.s[1:]
		return ret
	}
	return ""
}

func simplifyPath(path string) string {
	s := Stack{}
	s.Init()
	ret := ""
	var r int
	var tmp string
	path = path[1:]
	for len(path) > 0 {
		r = strings.IndexByte(path, '/')
		if r == -1 {
			if path == "." {

			} else if path == ".." {
				s.Pop()
			} else {
				s.Push(path)
			}
			break
		}
		if r == 0 {
			path = path[1:]
			continue
		}
		tmp = path[:r]
		path = path[r:]
		if tmp == "." {

		} else if tmp == ".." {
			s.Pop()
			continue
		} else {
			s.Push(tmp)
		}
	}
	for !s.Empty() {
		ret += "/" + s.LPop()
	}
	if len(ret) == 0 {
		ret = "/"
	}
	return ret
}
