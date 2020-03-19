package stack

// Stack is a stack
type Stack struct {
	Data []float32
}

func (s *Stack) push(x float32) {
	s.Data = append(s.Data, x)
}
