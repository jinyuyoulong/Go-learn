package main

import (
	"fmt"
)

type Human struct {
	Age   int
	Money int
	Look  int
	Sex   Sexual

	TargetLook int
	TargetSex  Sexual
}

type Sexual int

const (
	Male = iota
	Female
	Intersex
)

type MarryError struct {
	why string
}

func (me MarryError) Error() string {
	return me.why
}

func NewHuman(age, look, sex, money int, targetLook int, targetSex Sexual) *Human {
	h := new(Human)
	h.Age, h.Look, h.TargetLook = age, look, targetLook
	h.TargetSex = targetSex
	h.Money = money
	return h
}

func (h Human) Marray(o *Human) (happiness int, err error) {
	if h.TargetSex != o.Sex {
		panic("fuck off")
	}

	happiness = o.Look * o.Money / o.Age
	return
}
func main() {
	cook := NewHuman(23, 90, Male, 300, 80, Male)
	sister := NewHuman(33, 80, 200, Female, 99, Female)
	// h, e := cook.Marray(sister)
	h, e := sister.Marray(cook)
	if e != nil {
		fmt.Println(e.Error)
	}
	fmt.Printf("happiness is :%d\n", h)
}
