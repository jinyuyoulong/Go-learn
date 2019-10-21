package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeat string
	for index := 0; index < repeatCount; index++ {
		repeat += character
	}
	return repeat
}
