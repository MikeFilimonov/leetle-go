package leetstrings

func RomanToInt(input string) int {

	rates := make([]int, 89)
	rates[int('I')] = 1
	rates[int('V')] = 5
	rates[int('X')] = 10
	rates[int('L')] = 50
	rates[int('C')] = 100
	rates[int('D')] = 500
	rates[int('M')] = 1000

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
