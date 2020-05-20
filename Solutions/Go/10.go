package main

import (
	"fmt"
	"strings"
)

func main() {
	ret := isMatch("aa", ".*")
	// ret := strings.Split("abc", "")[1:3]

	fmt.Println(ret)
}

func isMatch(s string, p string) bool {
	// len
	if len(p) > len(s) {
		return false
	}
	// ''
	if s == "" && p == "" {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	// s,p => []
	// str := strings.Split(s, "")
	var ptr = 0
	pstr := strings.Split(p, "")
	for k, v := range pstr {
		if k == 0 && (v == "." || v == "*") {
			return false
		}
		last := ptr - 1
		if last < 0 {
			last = 0
		}
		if v == "." {
			if s[last] != s[ptr] {
				return false
			}
			ptr += 1
		} else if v == "*" {
			// abab/ab* ; ptr = 2
			step := len(s[:ptr])
			ptr += step
			if len(s) < ptr {
				return false
			}
			if s[:step] != s[step:2*step] {
				return false
			}
		} else {
			// abab/ab* ; ptr = 1
			if s[ptr:ptr+1] != v {
				return false
			}
			ptr += 1
		}

	}
	if len(s) <= ptr {
		return true
	}
	return false
}
