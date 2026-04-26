package leetarrays

func LargestArea(input []int) int {

	lefto := 0
	righto := len(input) - 1
	area := 0

	for lefto != righto {

		area = max(area,
			(righto-lefto)*min(input[lefto], input[righto]))

		if input[lefto] < input[righto] {
			lefto++
		} else {
			righto--
		}

	}

	return area

}

func bruteTrash(input []int) int {

	area := 0
	for i := 0; i < len(input)-1; i++ {

		for j := i + 1; j < len(input); j++ {
			area = max(min(input[i], input[j])*(j-i), area)
		}

	}

	return area

}
