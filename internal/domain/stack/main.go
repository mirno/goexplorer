package stack

type Stack []interface{} // []interface{} is a slice of the empty interface.

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]

	return item, true
}

func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}