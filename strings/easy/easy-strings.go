package leetstrings

func RomanToInt(input string) int {

	rates := make(map[rune]int)
	rates['I'] = 1
	rates['V'] = 5
	rates['X'] = 10
	rates['L'] = 50
	rates['C'] = 100
	rates['D'] = 500
	rates['M'] = 1000

	result := 0
	inputLength := len(input)

	for i := 0; i < inputLength; i++ {

		if i+1 < inputLength {

			if rates[rune(input[i])] < rates[rune(input[i+1])] {
				result += rates[rune(input[i+1])] - rates[rune(input[i])]
				i++
			} else {
				result += rates[rune(input[i])]
			}

		} else {
			result += rates[rune(input[i])]
		}

	}

	return result
}
