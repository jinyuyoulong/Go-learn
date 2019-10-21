package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("want '%d' but got '%d' ,given '%v'", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("测试多个array总和，汇总输出各项", func(t *testing.T) {
		number1 := []int{1, 2, 3}
		number2 := []int{1, 2}
		got := SumAll(number1, number2)
		want := []int{6, 3}

		// reflect.DeepEqual 该方法有点问题，不同类型的参数也可以编译通过，使用时比较简洁但需要注意
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want '%v' but got '%v'", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	// reflect.DeepEqual 该方法有点问题，不同类型的参数也可以编译通过，使用时比较简洁但需要注意
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want '%v' but got '%v'", want, got)
	}
}
