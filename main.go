package main

import (
	"container/list"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var legalArthmeticRegex string = `^(\d+(\.\d+)?|\+|\-|\*|\/|and|or|\(|\)|==|>=|<=|!=|>|<)+$`

func main() {
	test := "11+3*(2-4)/2"
	res, err := calculator(test)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result : ", res)
}

func calculator(data string) (float64, error) {
	match, err := regexp.MatchString(legalArthmeticRegex, data)
	if err != nil {
		return 0, err
	}
	if !match {
		return 0, errors.New("the data is not valid.")
	}
	tmpStr := suffixExpress(data)
	stack := list.New()
	for _, val := range tmpStr {
		if !isOperate(val[0]) {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0, err
			}
			stack.PushFront(f)
			continue
		}
		// a := stack.Front()
		// b := a.Next()

		if stack.Len() == 1 {
			if val == "+" || val == "-" {
				a := stack.Front()
				res := cal(0, a.Value.(float64), val)
				stack.Remove(a)
				stack.PushFront(res)
			}
		}
		if stack.Len() > 1 {
			a := stack.Front()
			b := a.Next()
			res := cal(b.Value.(float64), a.Value.(float64), val)
			stack.Remove(a)
			stack.Remove(b)
			stack.PushFront(res)
		}
	}
	if stack.Len() != 1 {
		return 0, errors.New("the result is no true.")
	}
	return stack.Front().Value.(float64), nil
}

func cal(a, b float64, s string) float64 {
	switch s {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func suffixExpress(data string) []string {
	stack := list.New()
	queue := list.New()
	leng := len(data)
	for i := 0; i < leng; {
		s := data[i]
		switch s > 0 {
		case s > 47 && s < 58:
			number := getNumber(data[i:])
			queue.PushBack(number)
			i = i + len(number)
		case isOperate(s):
			for e := stack.Front(); e != nil; {
				if !isPriority(s, e.Value.(byte)) {
					queue.PushBack(e.Value)
					b := e
					e = e.Next()
					stack.Remove(b)

				} else {
					e = e.Next()
				}
			}
			stack.PushFront(s)
			i++
		case s == '(':
			stack.PushFront(s)
			i++
		case s == ')':
			leng := stack.Len()
			for i := 0; i < leng; i++ {
				e := stack.Front()
				stack.Remove(e)
				if e.Value.(byte) == '(' {
					break
				}
				queue.PushBack(e.Value)
			}
			i++
		}
	}
	for e := stack.Front(); e != nil; e = e.Next() {
		queue.PushBack(e.Value)
	}
	suffix := make([]string, 0)
	for e := queue.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case string:
			suffix = append(suffix, e.Value.(string))
		case byte:
			suffix = append(suffix, string(e.Value.(byte)))
		}
	}
	return suffix
}

func isPriority(a, b byte) bool {
	if a == '*' || a == '/' {
		if b == '+' || b == '-' {
			return true
		}
		return false
	}

	return false
}

func isOperate(s byte) bool {
	switch s > 0 {
	case s == '+':
		return true
	case s == '-':
		return true
	case s == '*':
		return true
	case s == '/':
		return true
	default:
		return false
	}
	return false
}

func getNumber(str string) string {
	s := make([]rune, 0)
	for _, v := range str {
		if v < 48 || v > 57 {
			break
		}
		s = append(s, v)
	}
	return string(s)
}
