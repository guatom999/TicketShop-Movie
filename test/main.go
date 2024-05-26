package main

import "fmt"

func main() {
	output := []int{1, 2, 3, 4}
	result := make([]int, 0)

	for i, _ := range output {
		if i == 2 {
			result = append(result, 9)
		}
		result = append(result, output[i])
	}

	fmt.Println("result", result)

}
