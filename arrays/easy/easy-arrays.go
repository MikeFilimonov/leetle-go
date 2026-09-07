package leetarrays

func TwoSum(input []int, target int) []int {

	buffer := make(map[int]int)
	for k, v := range input {

		another := target - v
		if idx2, found := buffer[another]; found {
			return []int{idx2, k}
		}
		buffer[v] = k

	}

	return []int{}

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
