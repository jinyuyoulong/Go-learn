### 完整的测试驱动开发流程
- 编写测试
- 编写最少量的代码让测试运行并检查输出
- 编写足够的代码使测试通过
- 重构

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

### 未经检查的错误
虽然 Go 编译器对你有很大帮助，但有时你仍然会忽略一些事情，错误处理有时会很棘手。
有一种情况我们还没有测试过。要找到它，在一个终端中运行以下命令来安装 errcheck，这是许多可用的 linters（代码检测工具）之一。
go get -u github.com/kisielk/errcheck
然后，在你的代码目录中运行 errcheck .。
你应该会得到如下类似的内容：
wallet_test.go:17:18: wallet.Withdraw(Bitcoin(10))
这告诉我们的是，我们没有检查在代码行中返回的错误。我的计算机上的这行代码与我们的正常 withdraw 的场景相对应，因为我们没有检查 Withdraw 是否成功，因此没有返回错误。