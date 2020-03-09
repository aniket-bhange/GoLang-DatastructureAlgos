package sort

import swap "DS/utils"

//Bubble Sort sort array in Ologn^2
func Bubble(arr []int) []int {

	length := len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < (length-i)-1; j++ {
			if arr[j] > arr[j+1] {
				swap.Swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}
