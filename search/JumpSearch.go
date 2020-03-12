package search

import (
	"math"
)

//Jump searching methond for jumping around index and findind element
func Jump(arr []int, x int) int {

	length := len(arr)

	step := math.Floor(math.Sqrt(float64(length)))
	prev := 0

	min := math.Min(float64(step), float64(length))
	// fmt.Println("Steps are:", step, arr[int(min-1)])
	for arr[int(min-1)] < x {
		prev = int(step)
		step += math.Floor(math.Sqrt(float64(length)))
		// fmt.Println("in first loop:", prev, step, arr[int(min-1)])
		if length < prev {
			return -1
		}
		min = math.Min(float64(step), float64(length))
	}

	for arr[prev] < x {
		prev++
		if float64(prev) == math.Min(step, float64(length)) {
			return -1
		}
		prev = int(prev)
	}

	if arr[prev] == x {
		return prev
	}

	return -1

}
