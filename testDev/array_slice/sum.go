package main

func Sum(numbers []int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}

func SumAll(sliecs ...[]int) (sumall []int) {

	sliecLenth := len(sliecs)
	sumall = make([]int, sliecLenth)

	for i, msliec := range sliecs {
		sumall[i] = Sum(msliec)
		// sumall = append(sumall, isum)
	}
	return sumall
}
