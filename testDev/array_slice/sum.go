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

// SumAllTails 计算尾合集
func SumAllTails(sliecs ...[]int) (sums []int) {

	for _, msliec := range sliecs {
		var tail []int
		if len(msliec) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail = msliec[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
