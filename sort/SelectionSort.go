package sort

import swap "DS/utils"

//Selection sort is sorting technique with Olog(n^2)
func Selection(arr []int) []int {

	val := arr
	length := len(val)
	for i := 0; i < length; i++ {
		minIndx := i

		for j := (i + 1); j < length; j++ {
			if val[minIndx] > val[j] {
				minIndx = j
			}
		}
		swap.Swap(&val[i], &val[minIndx])
		// val[i], val[minIndx] = val[minIndx], val[i]
	}
	return val
}
