package main

import (
	"DS/sort"
	"fmt"
	"os"
	"strconv"
)

//ConvertStoI take the string array and retuen int array
func ConvertStoI(strArr []string) ([]int, error) {
	var op []int
	for _, v := range strArr {
		j, _ := strconv.Atoi(v)
		op = append(op, j)
	}
	return op, nil
}

func main() {
	// val := []int{2, 5, 7, 8, 9, 4}
	var input []string
	input = os.Args[1:]

	inputValue, err := ConvertStoI(input)

	if err != nil {
		fmt.Println(err)
	}

	// inputValue = sort.Selection(inputValue)
	// inputValue = sort.Bubble(inputValue)
	// inputValue = sort.MergeSort(inputValue)
	inputValue = sort.HeapSort(inputValue)

	fmt.Println(inputValue)
}
