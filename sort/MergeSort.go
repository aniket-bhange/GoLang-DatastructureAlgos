package sort

//Merge is method to merge two arrays
func Merge(leftArr []int, rightArr []int) (result []int) {

	result = make([]int, len(leftArr)+len(rightArr))
	i := 0
	for len(leftArr) > 0 && len(rightArr) > 0 {
		if leftArr[0] < rightArr[0] {
			result[i] = leftArr[0]
			leftArr = leftArr[1:]
		} else {
			result[i] = rightArr[0]
			rightArr = rightArr[1:]
		}
		i++
	}

	for j := 0; j < len(leftArr); j++ {
		result[i] = leftArr[j]
		i++
	}
	for j := 0; j < len(rightArr); j++ {
		result[i] = rightArr[j]
		i++
	}

	return
}

//MergeSort function break array into half and then sort every element and merge it again
//Complexity is O(nlog n)
func MergeSort(arr []int) []int {

	length := len(arr)
	if length == 1 {
		return arr
	}

	mid := length / 2

	leftArr := arr[:mid]
	rightArr := arr[mid:]

	return Merge(MergeSort(leftArr), MergeSort(rightArr))
}
