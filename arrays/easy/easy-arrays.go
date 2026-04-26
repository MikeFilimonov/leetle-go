package arrays

import (
	"fmt"
	"slices"
)

func TwoSum() {

	input := []int{2, 7, 11, 15}

	fmt.Println("the indices are: ", twoSum(input, 9))

}

func twoSum(input []int, target int) []int {

	result := make([]int, 0)

	for i, v := range input {

		nextValue := target - v
		coopIndex := slices.Index(input[i+1:], nextValue)

		if coopIndex == -1 {
			continue
		} else {

			result = append(result, i)
			result = append(result, i+coopIndex+1)
			break
		}

	}

	return result

}
