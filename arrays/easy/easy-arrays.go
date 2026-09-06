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

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	shorty := strs[0]
	for _, v := range strs {
		if len(v) < len(shorty) {
			shorty = v
		}
	}

	for i := 0; i < len(shorty); i++ {

		starter := strs[0][i]

		for _, s := range strs[1:] {
			if s[i] != starter {
				return strs[0][:i]
			}
		}

	}

	return shorty

}

func ContainsDuplicate(nums []int) bool {

	cloneBuster := make(map[int]struct{})
	for _, v := range nums {

		if _, found := cloneBuster[v]; found {
			return true
		}
		cloneBuster[v] = struct{}{}
	}

	return false

}
