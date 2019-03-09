package main

import (
	"fmt"
	"math"
)

func main() {

	var s = "abcabcbb"

	//fmt.Println(bruteForce(s))
	fmt.Println(slidingWindow(s))
	//fmt.Println(slidingWindowOptimized(s))
}

func bruteForce(s string) int {

	return 0
}

func slidingWindow(s string) int {
	var n = len(s)
	var m = make(map[string]int)
	var i, j int
	var max float64

	for i < n && j < n {
		if _, ok := m[string(s[j])]; !ok {
			m[string(s[j])] = 1
			j++
			max = math.Max(max, float64(j-i))
		} else {
			delete(m, string(s[i]))
			i++
		}
	}

	return int(max)
}

func slidingWindowOptimized(s string) int {

	return 0
}
