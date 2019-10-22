测试命令 `go test`

辅助函数
```
// t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）

assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
t.Run("say hello to people", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
```

显示详细测试信息
go test -v

// ExampleAdd 示例代码
// 必须要有Output 这行注释
```
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```


// BenchmarkRepeat 基准测试
// go test -bench=.
```
// goos: darwin
// goarch: amd64
// pkg: github.com/jinyuyoulong/Go-learn/testDev/iteration
// BenchmarkRepeat-4        6746726               173 ns/op
// PASS
// ok      github.com/jinyuyoulong/Go-learn/testDev/iteration      1.354s

//  说明运行一次这个函数需要 173 纳秒
// 注意：基准测试默认是顺序运行的。
```
覆盖率命令
go test -cover

// 使用类型别名具有可描述性
type Bitcoin int