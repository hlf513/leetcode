package main

import (
	"fmt"
	"strings"
)

func main() {
	//var path = "/a/../../b/../c//.//"
	//var path = "/home//foo/"
	//var path = "/a//b////c/d//././/.."
	var path = "/../"

	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	n := len(paths)

	var depth = 0
	var res = ""
	for i := n - 1; i >= 0; i-- {
		if paths[i] == "" || paths[i] == "." {
			continue
		}
		if paths[i] == ".." {
			depth++
			continue
		}
		if depth > 0 {
			depth--
			continue
		}
		res = "/" + paths[i] + res
	}

	if res == "" {
		res = "/"
	}

	return res
}
