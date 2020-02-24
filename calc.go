package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type IntStack struct {
	stack []int
}

func (st *IntStack) Push(elem int) bool {
	st.stack = append(st.stack, elem)
	return true
}

func (st *IntStack) Pop() int {
	if (len(st.stack)) < 1 {
		return 0
	}
	res := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	return res
}
type Stack struct {
	stack []string
}

func (st *Stack) Push(elem string) bool {
	st.stack = append(st.stack, elem)
	return true
}

func (st *Stack) Pop() string {
	if (len(st.stack)) < 1 {
		return ""
	}
	res := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	return res
}

func (st Stack) Check() string {
	if (len(st.stack)) < 1 {
		return ""
	}
	return st.stack[len(st.stack)-1]
}

func (st Stack) Len() int {
	return len(st.stack)
}

func CheckLast(last *string, resStr *string) {
	if *last != "" {
		*resStr += *last+" "
		*last = ""
	}
}

func skipSpaces(s []byte) []byte {
	c, w := utf8.DecodeRune(s)
	for w > 0 && unicode.IsSpace(c) {
		s = s[w:]
		c, w = utf8.DecodeRune(s)
	}
	return s
}

func readDigits(s []byte) (numStr, remain []byte) {
	numStr = s
	totalW := 0
	c, w := utf8.DecodeRune(s)
	for w > 0 && unicode.IsDigit(c) {
		s = s[w:]
		totalW += w
		c, w = utf8.DecodeRune(s)
	}
	return numStr[:totalW], s
}

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func Evaluate(s []byte) int {
	stack := IntStack{}
	var a, b int
	var token []byte

	s = skipSpaces(s)
	for len(s) > 0 {
		c, w := utf8.DecodeRune(s)
		switch {
		case unicode.IsDigit(c):
			token, s = readDigits(s)
			num, err := strconv.Atoi(string(token))
			if err != nil {
				fmt.Println(err)
			} else {
				stack.Push(num)
			}
		case c == '+':
			b = stack.Pop()
			a = stack.Pop()
			stack.Push(a+b)
			s = s[w:]
		case c == '-':
			b = stack.Pop()
			a = stack.Pop()
			stack.Push(a-b)
			s = s[w:]
		case c == '*':
			b = stack.Pop()
			a = stack.Pop()
			stack.Push(a*b)
			s = s[w:]
		case c == '/':
			b = stack.Pop()
			a = stack.Pop()
			if b == 0 {
				fmt.Println("zero division error")
				os.Exit(1)
			}
			stack.Push(a/b)
			s = s[w:]
		}
		s = skipSpaces(s)
	}
	return stack.Pop()
}

func Tokenize(str string) string {
	stack := Stack{}
	resStr := ""
	last := ""
	for _, elem := range str {
		if unicode.IsDigit(elem) {
			last += string(elem)
		} else {
			CheckLast(&last, &resStr)
			if elem == '-' || elem == '+' {
				if stack.Len() > 0 && stack.Check() != "(" {
					resStr += stack.Pop() + " "
				}
			}
			if elem == '*' || elem == '/' {
				if stack.Check() == "/" || stack.Check() == "*" {
					resStr += stack.Pop() + " "
				}
			}

			if elem == ')' {
				CheckLast(&last, &resStr)
				for stack.Check() != "(" {
					resStr += stack.Pop() + " "
				}
				stack.Pop()
			} else {
				stack.Push(string(elem))
			}
		}
	}
	CheckLast(&last, &resStr)
	for stack.Len() > 0 {
		resStr += stack.Pop() + " "
	}
	return resStr
}

func CheckSigns(runa rune) bool {
	if runa == '+' || runa == '-' || runa == '/' || runa == '*' {
		return true
	}
	return false
}

func Validation(query string) bool {
	if CheckSigns(rune(query[0])) || CheckSigns(rune(query[len(query)-1])) {
		return false
	}
	for i, runa := range query {
		if unicode.IsDigit(runa) || runa == '(' || runa == ')' || CheckSigns(runa) {
			if (i > 0) && CheckSigns(runa) && (CheckSigns(rune(query[i-1]))) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func CheckBrackets(query string) bool {
	b_stack := Stack{}
	for _, runa := range query {
		if runa == '(' {
			b_stack.Push(string(runa))
			continue
		}
		if runa == ')' {
			if b_stack.Check() == "(" {
				b_stack.Pop()
				continue
			} else {
				return false
			}
		}
	}
	if b_stack.Len() > 0 {
		return false
	}
	return true
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("not enough arguments to work with")
		os.Exit(1)
	}
	query := os.Args[1]
	if !Validation(query) {
		fmt.Println("error with data")
		os.Exit(1)
	}
	if !CheckBrackets(query) {
		fmt.Println("error with brackets")
		os.Exit(1)
	}
	fmt.Println(Evaluate([]byte(Tokenize(query))))
}