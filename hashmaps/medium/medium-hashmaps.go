package leethashmaps

import (
	"strings"
)

func intToRoman(num int) string {

	arabians := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var sb strings.Builder
	focus := 0
	for num > 0 {

		if arabians[focus] > num {
			focus++
			continue
		}
		num -= arabians[focus]
		sb.WriteString(romans[focus])
	}

	return sb.String()

}
