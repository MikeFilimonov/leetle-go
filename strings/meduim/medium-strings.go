package leetstrings

func LengthOfLongestSubstring(input string) int {

	if len(input) <= 1 {
		return len(input)
	}

	validator := make(map[rune]int)
	result := 0
	l := 0
	for r := 0; r < len(input); r++ {

		if value, found := validator[rune(input[r])]; found {
			l = max(l, value+1)
		}

		validator[rune(input[r])] = r
		result = max(result, r-l+1)

	}

	return result

}
