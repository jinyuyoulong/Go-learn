[TOC]

# 为什么要用 gofmt

大多数人认为我们格式化Go代码gofmt 以使代码看起来更好或结束团队成员之间关于程序布局的争论。但是， [1](https://blog.v5u.win/post/go-gofmt/#fn:gofmt)是，如果一个算法定义围棋源代码是如何被格式化，然后编程，如goimports 或 gorename 或者 go fix，可以更容易地编辑源代码，而无需编写代码回来时引入伪格式更改。这有助于您长期维护代码。

------

> 我认为源代码格式化工具非常简洁 提供。解析器/ ast /漂亮也是一件好事 打印机可用作mod。
>
> 但是，我不确定是否尝试强制执行格式化样式 拒绝配置格式化程序的能力（如FAQ提及） 是个好主意。至少go /打印机包可能是好的 即使命令行gofmt工具，也允许更灵活的配置 没有。 对于gofmt的输出，没有人或将永远不会满意， 但它实际上相当不错，更重要的是，人们适应 令人惊讶的是，这种风格起初看起来很陌生。特别 在格式化方面，风格实际上只是“你习惯的”。 由于Go是一种全新的语言，所以不应该这样 很难习惯不同的格式。

我们希望人们能够准确地接受gofmt的输出 因为它结束了这种风格的辩论。怎么样 C有很多不同的支撑款式吗？太多。

就个人而言，我发现让gofmt格式化为我是一种解放， 因为这意味着我有更多可用的神经元 攻击有趣的编程问题。有 我不喜欢gofmt输出的东西，但我不喜欢 不再担心他们了。

但所有这些都错过了我认为最激动人心的事情 关于gofmt：事实上我们有一个可以拿起每个人的工具 Go树中的源文件，将其解析为内部表示， 然后将完全相同的字节放回原位。 （很大一部分 这是由于进入gofmt的工作量，以及 其余的是因为我们同意对gofmt进行标准化 一旦你拥有了这样一个工具，它变得非常容易 在中间插入机械处理，解析之间 和印刷。所以我们拥有一个程序的所有难点 操纵工具只是坐着等待使用。我已经开始了 以前为C编写类似的工具，从来没有得到过输出 完全匹配输入。同意接受“gofmt风格” 是在有限数量的代码中使它可行的部分。

我希望人们会按照自己的代码使用gofmt。 正如我所说，它不需要超过几个星期 习惯于新的编码风格，特别是如果你是一个品牌 新语言，每个人都有巨大的利益 相同的风格。

我们当然打算继续格式化所有代码 使用gofmt去树。像Go一样，这是一个实验。

Russ

------

1. 下面有邮件组的讨论过程 [[return\]](https://blog.v5u.win/post/go-gofmt/#fnref:gofmt)

