package main

import "fmt"

func main() {
	grid := [][]int{
		[]int{1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	var hashCount = make(map[string]int)
	var next = make(map[string]string)
	var max = 0

	for k, v := range grid {
		for k1, v1 := range v {
			if v1 == 1 {
				nextKey := string(k) + "-" + string(k1)
				if _, ok := next[nextKey]; ok {
					// 找到头节点
					i := nextKey
					for next[i] != "-1" {
						i = next[i]
					}
					hashCount[i]++
				} else {
					if k1 == 0 { // 判断上、下
						up := string(k-1) + "-" + string(k1)
						if _, ok := next[up]; ok {

						}
					} else { // 判断左、上、下

					}
				}
			} //end if
		}
	}

	for _, v := range hashCount {
		if v > max {
			max = v
		}
	}

	return max
}
