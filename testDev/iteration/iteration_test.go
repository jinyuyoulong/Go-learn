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

// BenchmarkRepeat 基准测试
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/jinyuyoulong/Go-learn/testDev/iteration
// BenchmarkRepeat-4        6746726               173 ns/op
// PASS
// ok      github.com/jinyuyoulong/Go-learn/testDev/iteration      1.354s

//  说明运行一次这个函数需要 173 纳秒
// 注意：基准测试默认是顺序运行的。
