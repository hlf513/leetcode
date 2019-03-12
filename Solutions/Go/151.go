package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "the sky  is blue"
	//var s =" "

	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	tmp := strings.Split(strings.Trim(s, " "), " ")
	n := len(tmp)

	var r []string
	for i := n - 1; i >= 0; i-- {
		if tmp[i] != "" {
			r = append(r, tmp[i])
		}
	}

	return strings.Join(r, " ")
}
