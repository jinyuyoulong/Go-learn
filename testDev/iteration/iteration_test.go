package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expeated := "aaaaa"
	if repeated != expeated {
		t.Errorf("expeated '%s' but got '%s'", expeated, repeated)
	}
}
