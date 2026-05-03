package leetstrings

func lengthOfLongestSubstring(input string) int {

	if len(input) <= 1 {
		return len(input)
	}

	validator := make(map[rune]int)
	result := 0
	// for i := 0; i < len(input); {

	// 	input = input[i:]
	// 	for k, v := range input {

	// 		if previous, found := validator[v]; found {

	// 			validator[v] = k
	// 			result = max(k, result)
	// 			i = previous + 1
	// 			break

	// 		}

	// 		validator[v] = k

	// 	}

	// }

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
