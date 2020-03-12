package main

import (
	"DS/search"
	"fmt"
	"os"
	"strconv"
)

//ConvertStoI take the string array and retuen int array
func ConvertStoI(strArr []string) ([]int, error) {
	var op []int
	for _, v := range strArr {
		if v == "-f" {
			return op, nil
		}
		j, _ := strconv.Atoi(v)
		op = append(op, j)
	}
	return op, nil
}

func main() {
	// val := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var input []string
	input = os.Args[1:]

	inputValue, err := ConvertStoI(input)

	if err != nil {
		fmt.Println(err)
	}

	valueToFind := os.Args[len(inputValue)+1:]
	find, _ := strconv.Atoi(valueToFind[1])
	// inputValue = sort.Selection(inputValue)
	// inputValue = sort.Bubble(inputValue)
	// inputValue = sort.MergeSort(inputValue)
	// inputValue = sort.HeapSort(inputValue)

	searchValue := search.Jump(inputValue, find)

	fmt.Println(inputValue)
	fmt.Println(searchValue)

}
