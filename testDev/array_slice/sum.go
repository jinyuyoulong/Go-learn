package main

func Sum(numbers []int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}
