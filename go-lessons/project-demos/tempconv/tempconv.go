// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroK Kelvin = -273.15
	FreezingK     Kelvin = 0
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g° C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g° F", f)
}

// 包管理 demo
