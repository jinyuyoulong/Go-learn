测试命令 `go test`

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