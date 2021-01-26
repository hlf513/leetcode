package main

import "fmt"

func main() {
	intervals := []Interval{
		{2, 6},
		{1, 3},
		{8, 10},
		{15, 18},
	}

	//intervals := []Interval{
	//	{1, 4},
	//	{1, 5},
	//}

	//intervals := []Interval{
	//	{1, 4},
	//	{0, 4},
	//}

	//intervals := []Interval{
	//	{1, 4},
	//	{2, 3},
	//}

	//intervals := []Interval{
	//	{1, 4},
	//	{0, 5},
	//}

	fmt.Println(merge(intervals))
}

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 1 {
		return []Interval{}
	}
	// 按照 start 排序
	quickSort(intervals)
	fmt.Println(intervals)

	// 开启合并区间
	var res []Interval
	res = append(res, intervals[0])
	intervals = intervals[1:]

	for i := 0; i < len(intervals); i++ {
		n := len(res)
		if intervals[i].Start <= res[n-1].End {
			if res[n-1].End < intervals[i].End { // 判断 end 值大小
				res[n-1].End = intervals[i].End
			}
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

func quickSort(intervals []Interval) {
	// 退出条件
	if len(intervals) <= 1 {
		return
	}

	// 处理
	mid := intervals[0].Start
	head, tail := 0, len(intervals)-1

	for i := 1; i <= tail; {
		if intervals[i].Start > mid {
			intervals[tail], intervals[i] = intervals[i], intervals[tail]
			tail--
		} else {
			// 相等,需再次判断 end 值
			if (intervals[i].Start == mid && intervals[i].End <= intervals[0].End) || intervals[i].Start < mid {
				intervals[head], intervals[i] = intervals[i], intervals[head]
			}
			head++
			i++
		}
	}

	// 下沉
	quickSort(intervals[head+1:])
	quickSort(intervals[:head])
}
