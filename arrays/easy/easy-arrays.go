package leetarrays

import (
	"slices"
	"strings"
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

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	shorty := strs[0]
	for _, v := range strs {
		if len(v) < len(shorty) {
			shorty = v
		}
	}

	var sb strings.Builder

	for i := 0; i < len(shorty); i++ {

		starter := strs[0][i]

		isCommonPrefix := true
		for c := 1; c < len(strs); c++ {
			isCommonPrefix = isCommonPrefix && (starter == strs[c][i])
			if !isCommonPrefix {
				break
			}
		}

		if !isCommonPrefix {
			break
		}
		sb.WriteByte(starter)

	}

	return sb.String()

}
