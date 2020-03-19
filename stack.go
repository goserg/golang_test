package stack

// Stack is a stack
type Stack struct {
	Data []float32
}

// Push appends x to Data
func (s *Stack) Push(x float32) {
	s.Data = append(s.Data, x)
}

// Peek returns x from top of the stack
func (s *Stack) Peek() float32 {
	return s.Data[len(s.Data)-1]
}
