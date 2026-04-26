package leetarrays

import (
	"slices"
)

func TwoSum(input []int, target int) []int {

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
