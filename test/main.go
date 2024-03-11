package main

import "fmt"

func main() {

	test := []int{1, 2, 3, 4}

	EachCons(test, 3)

}

func EachCons(arr []int, n int) [][]int {
	// your code here
	result := make([][]int, 0)
	for i, _ := range arr {
		if (i + n) > len(arr) {
			break
		}
		result = append(result, arr[i:i+n])
	}

	fmt.Println("print res", result)

	return result
}
