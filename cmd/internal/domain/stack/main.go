package stack

type Stack []interface{}

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