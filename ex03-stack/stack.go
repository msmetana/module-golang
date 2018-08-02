package stack

type Stack struct {
	stack []int
	pos   int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(elem int) {
	s.pos += 1
	s.stack = append(s.stack, elem)
}

func (s *Stack) Pop() int {
	s.pos -= 1

	return s.stack[s.pos]
}
