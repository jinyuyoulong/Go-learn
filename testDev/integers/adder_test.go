// adder_test
package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// ExampleAdd 示例代码
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// 必须要有Output 这行注释

// === RUN   TestAddr
// --- PASS: TestAddr (0.00s)
// === RUN   ExampleAdd
// --- FAIL: ExampleAdd (0.00s)
// got:
// 3
// want:
// 6
// FAIL
// exit status 1
