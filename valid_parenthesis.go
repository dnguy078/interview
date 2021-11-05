package main

func isValid(s string) bool {
	var validLetters = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := Stack{data: []string{}}
	for _, c := range s {
		if isClosingParenthesis(string(c)) {
			top := stack.Peek()
			if complement, valid := validLetters[top]; !valid || complement != string(c) {
				return false
			}

			stack.Pop()
		} else if isOpenParenthesis(string(c)) {
			stack.Push(string(c))
		}
	}
	if !stack.IsEmpty() {
		return false
	}
	return true
}

func isValidParenthesis(s string) bool {
	var stack []string
	var corresponding = map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	for i := 0; i < len(s); i++ {
		charAt := s[i]
		if openBracket, isClosing := corresponding[string(charAt)]; isClosing {
			// pop from stack
			top := stack[len(stack)-1]
			if top == openBracket {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(charAt))
		}
	}
	return len(stack) == 0
}

type Stack struct {
	data []string
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return s.data[len(s.data)-1]
}
func (s *Stack) Push(input string) {
	s.data = append(s.data, input)
}

func (s *Stack) Pop() string {
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element
}

func isClosingParenthesis(c string) bool {
	return c == ")" || c == "]" || c == "}"
}

func isOpenParenthesis(c string) bool {
	return c == "(" || c == "[" || c == "{"
}
