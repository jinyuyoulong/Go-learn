run.sh
========
当前目录下执行 go test

注意，你不必在多个测试框架之间进行选择，然后再去理解如何安装它们。你需要的一切都内建在 Go 语言中，语法与你将要编写的其余代码相同。
编写测试
编写测试和函数很类似，其中有一些规则
程序需要在一个名为 xxx_test.go 的文件中编写
测试函数的命名必须以单词 Test 开始
测试函数只接受一个参数 t *testing.T
现在这些信息足以让我们明白，类型为 *testing.T 的变量 t 是你在测试框架中的 hook（钩子），所以当你想让测试失败时可以执行 t.Fail() 之类的操作。

参考 https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/hello-world