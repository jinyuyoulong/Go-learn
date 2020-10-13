本文摘自 https://www.kancloud.cn/cserli/golang/524388

本文译自go-proverbs, 脱胎于 Rob Pike 振奋人心的演讲视频 talk at Gopherfest SV 2015 (bilibili).

**不要通过共享内存进行通信, 通过通信共享内存 (Don’t communicate by sharing memory, share memory by communicating)**

传统的线程模型（通常在编写 Java, C++ 和 Python 程序时使用）要求程序员使用共享内存在线程之间进行通信. 通常, 共享数据结构受锁保护, 线程将争夺这些锁访问数据, 在某些情况下, 通过使用 Python 的 Queue 等线程安全的数据结构可以使这变得更容易.

Go 的并发原语 (goroutines 和 channels) 为构造并发软件提供了一种优雅而独特的手段. (这些概念有一个有趣的历史, 要从 C.A.R.Hoare 的通信顺序进程说起.) Go 鼓励使用 channels 在 goroutines 之间传递对数据的引用, 而不是显式地使用锁来调解对共享数据的访问. 这种方法确保只有一个 goroutine 可以在给定的时间访问数据. 这个概念总结在 Effective Go 文档中 (任何 Go 程序员都必须阅读).

Go 官方博客中有一篇文章对该谚语解读, 可以参见原文.

**并发不是并行 (Concurrency is not parallelism)**

当人们听到 并发 这个词的时候, 他们经常会想到并行, 这是一个相关的, 但非常独特的概念. 在编程中, 并发是独立执行的进程的组成, 而并行则是 (可能相关的) 计算的同时执行. 并发是一次处理很多事情. 并行是一次做很多事情.

Channels 重排序; 互斥量串行化 (**Channels orchestrate; mutexes serialize**)

这个看中文（翻译待商榷）是不是一脸懵 (虽然英文也看不懂) ? 其实分号前后说的是一个意思, 该谚语按我的个人理解可以用 go 程序 (来自 go tour) 解释成如下:

```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // 此处如果改成互斥量一样可以做到
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c

    fmt.Println(x, y, x+y)
}
```

**接口越大, 抽象越弱 (The bigger the interface, the weaker the abstraction)**

接口背后的概念是通过将对象的行为抽象为简单的契约来允许重用性. 虽然接口不是 Go 专有的, 但由于 Go 接口通常趋向于小型化, Go 程序员才广泛使用它们. 通常情况下, 一个接口只限于一到两个方法.

Go io 包接口就是典型的例子.

**充分利用零值 (Make the zero value useful)**

零值的典型例子如 bytes.Buffer 和 sync.Mutex:

```go
var buf bytes.Buffer
buf.Write([]byte("hello"))
fmt.Println(buf.String())

var mu sync.Mutex
mu.Lock()
mu.Unlock()
```

这样看起来是不是感觉一点用没有 ? 如果这样呢 ?

```go
type Map struct {
    mu sync.RWMutex
    // ...
}

func (m *Map) Set(k, v interface{}) {
    m.mu.Lock()
    defer m.mu.Unlock()
    if m.m == nil {
        m.m = make(map[interface{}]interface{})
    }
    m.m[k] = v
}
```

**interface{} 言之无物 (interface{} says nothing)**

该谚语不是说 interface {} 不代表任何东西, 而是说该类型无静态检查以及调用时保证, 比如你的 func 接收一个 interface{} 类型, 你写的时候是可用的, 但是某个时间你进行了代码重构可能坏掉了.

**Gofmt 的风格没有人喜欢, 但是 gofmt 是每个人的最爱 (Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite)**

该谚语告诉我们少些风格之争, 用这些时间多写代码.

**小复制好过小依赖 (A little copying is better than a little dependency)**

简单说就是如果你可以手动撸小快代码就不要导入一个库去做, 比如 UUID:

```go
// see: https://groups.google.com/d/msg/golang-nuts/d0nF_k4dSx4/rPGgfXv6QCoJ
package main

import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("/dev/urandom") // 演示用忽略 errcheck
    b := make([]byte, 16)
    f.Read(b)
    f.Close()
    uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    fmt.Println(uuid)
}
```

虽然有一堆写好的 UUID 库, 当你仅仅需要一个 UUID v4 实现.

**系统调用必须始终使用构建标签保证 (Syscall must always be guarded with build tags)**

不同的系统 (*NIX, Windows) 调用导致你同一个 func (实现并不一样) 可能需要在不同的系统上构建才能得到你想要的结果. 简单说就是系统调用不可移植才这么干. 示例可参见 Go 标准库 syscall.

**Cgo 必须始终使用构建标签保证 (Cgo must always be guarded with build tags)**

基本上原因同上一条.

**Cgo 不是 Go (Cgo is not Go)**

如果可能不要用 Cgo. 这里有篇文章说明了为什么.

**unsafe 包无保证 (With the unsafe package there are no guarantees)**

包如其名, 不安全. 你可以使用 unsafe 包如果你准备好了有一天它会坏掉.

**清晰好过聪明 (Clear is better than clever)**

Rob Pike 在他与别人合著的 <程序设计实践> 中写到: “写清晰的代码, 不要写聪明的代码”.

**反射永远不是清晰的 (Reflection is never clear)**

很多人在 Stackoverflow 上抱怨 Go 的反射不工作, 因为那不是为你准备的😂! 只有很少很少的人应该用反射这个非常强大而又非常难的特性. 新手应该远离反射和 interface{}.

**错误也是一种值 (Errors are values)**

值可以被编程, 并且由于错误是值, 所以错误可以被编程. Go 官方博客有对此的解读.

**不要止步于检查错误而要优雅的处理 (Don’t just check errors, handle them gracefully)**

Dave Cheney 有篇博客详细解读了该谚语.

**设计架构, 命名组件, 记录细节 (Design the architecture, name the components, document the details)**

当你写一个大型系统的时候, 你把它设计成一种结构化的东西. 想象组件的每一个部分并行工作, 为不同的组件起好的名字, 因为这些名字会出现在稿纸上.

拿 Go 程序来说, 如果名字不错, 组件就好理解, 那么程序的结构设计就会清晰, 程序会感觉很自然.

但是还有很多东西你需要解释, 所以这些是你需要解释的细节. 但是命名会帮助你解释很大一部分设计. 细节只是填补材料的缺口可能用来为用户打印工程图解文档.

**文档是针对用户的 (Documentation is for users)**

很多人写文档表明某个 func 是做什么的, 但是他们不想想这个 func 是为谁而写. 这有很大的不同. 你知道这个 func 返回什么是对的, 但是它为什么返回了你使用的时候不一样的结果?

把自己当成使用者而不是写它的人, 那么 godoc 上的文档就是对用户有用的. 这对于其他语言一样适用.

**不要慌 (Don’t panic)**

不要使用 panic 进行正常的错误处理. 使用错误 (error) 和多个返回值.