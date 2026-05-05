package main

import (
	"fmt"
	easyArrays "leetle-go/arrays/easy"
	mediumArrays "leetle-go/arrays/medium"
	easyStrings "leetle-go/strings/easy"
	mediumStrings "leetle-go/strings/meduim"
)

const defaultTemplate = "Input: %v, output: %v.\r\n"

func main() {

	easyRunner()
	mediumRunner()

}

func easyRunner() {

	data := []int{2, 7, 11, 15}
	target := 9
	nums := easyArrays.TwoSum(data, target)
	fmt.Printf(defaultTemplate, data, nums)

	fmt.Println("Roman to int")
	input := "MCMXCIV"
	result := easyStrings.RomanToInt(input)
	fmt.Printf(defaultTemplate, input, result)
}

func mediumRunner() {

	fmt.Println("container with most water")
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := mediumArrays.LargestArea(input)
	fmt.Printf(defaultTemplate, input, result)

	fmt.Println("length of the longest substring")
	initialData := "abbjahdbls"
	result = mediumStrings.LengthOfLongestSubstring(initialData)
	fmt.Printf(defaultTemplate, input, result)
}
