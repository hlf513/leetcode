package main

import (
	"fmt"
)

func main() {
	//s := "()[]{}"
	s := "(]"
	fmt.Println(isValid20(s))
}

func isValid20(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if string(s[i]) == ")" && pop != "(" {
				return false
			}
			if string(s[i]) == "]" && pop != "[" {
				return false
			}
			if string(s[i]) == "}" && pop != "{" {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}

	return true
}
