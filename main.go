package main

import (
	"fmt"
	easyArrays "leetle-go/arrays/easy"
	mediumArrays "leetle-go/arrays/medium"
)

func main() {

	easyRunner()
	mediumRunner()

}

func easyRunner() {
	easyArrays.TwoSum()
}

func mediumRunner() {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := mediumArrays.LargestArea(input)
	fmt.Printf("Input: %v. Max area for the result: %v\n", input, result)
}
