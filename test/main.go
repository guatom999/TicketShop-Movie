package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// func main() {
// 	output := []int{1, 2, 3, 4}
// 	result := make([]int, 0)

// 	for i, _ := range output {
// 		if i == 2 {
// 			result = append(result, 9)
// 		}
// 		result = append(result, output[i])
// 	}

// 	fmt.Println("result", result)

// }

func main() {
	byteString := make([]byte, 4)

	_, err := rand.Read(byteString)
	if err != nil {
		fmt.Println("err", err)
	}

	randomString := hex.EncodeToString(byteString)

	fmt.Println("randomString ::::::::::>", randomString[:7])
}
