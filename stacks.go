package stacks

// Stack is an array of nil interfaces
type Stack []interface{}

// Push an element to the stack
func (s *Stack) Push(elem interface{}) {
	temp := &Stack{elem}
	*s = append(*temp, *s...)
}

// Pop an element and return it
func (s *Stack) Pop() interface{} {
	if len(*s) > 0 {
		elem := (*s)[0]
		*s = (*s)[1:]
		return elem
	}
	return nil
}

// ToString print the stack as a string
func (s *Stack) ToString() string {
	sString := "STACK TOP"
	sCopy := &Stack{}
	*sCopy = *s

	elem := sCopy.Pop()

	for elem != nil {
		sString = sString + "\n" + elem.(string)
		elem = sCopy.Pop()
	}

	return sString + "\nSTACK BOTTOM"
}
