package main

func Sum(numbers [5]int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}
